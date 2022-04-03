package main

import (
	"log"
	"math"
	"runtime"
	"sync"
	"time"

	. "github.com/tntmeijs/gengo/mathematics"
	. "github.com/tntmeijs/gengo/scene"
	. "github.com/tntmeijs/gengo/utility"
)

// =============================================================================================================================
// =============================================================================================================================
// =============================================================================================================================

// System
const GoRoutinesPerAvailableCpu = 1
const AverageNumberOfTasksPerWorker = 2

// Output image information
const ImageResolutionX = 640
const ImageResolutionY = 360
const ImageFileName = "output.png"

// Camera constants
const RayStepSize = 0.01
const CameraNearPlane = 0.001
const CameraFarPlane = 25.0

// Light parameters
const AmbientStrength = 0.25
const SpecularStrength = 0.5
const SpecularShininess = 32

// Camera transformation
var cameraPosition = Vec3{X: 0.0, Y: 0, Z: -1.525}
var cameraLookAt = Vec3{X: 0.0, Y: 0.0, Z: 0.0}

// Lights
var ambientLightPosition = Vec3{X: -10.0, Y: 10.0, Z: -10.0}

// Colors
var ambientColor = Color{Red: 255, Green: 255, Blue: 255, Alpha: 255}
var surfaceColor = Color{Red: 26, Green: 188, Blue: 156, Alpha: 255}

// =============================================================================================================================
// =============================================================================================================================
// =============================================================================================================================

var scene = NewScene(sceneSDF)
var camera = NewCamera(cameraPosition, cameraLookAt, CameraNearPlane, CameraFarPlane)

// =============================================================================================================================
// =============================================================================================================================
// =============================================================================================================================

// Represents a render task
type RenderTask struct {
	StartRow, RowCount int
}

// Represents the result of a render task
type RenderResult struct {
	StartRow, RowCount int
	Pixels             []Color
}

// Log the time since the start time
func trackTime(start time.Time, name string) {
	log.Printf("%s took %s", name, time.Since(start))
}

// Entire scene represented as a signed distance function
func sceneSDF(point Vec3) float64 {
	return MandelbulbSDF(point, 10, 8, 5.0)
}

// Simple Blinn-Phong lighting model
//
// Reference: https://learnopengl.com/Advanced-Lighting/Advanced-Lighting
func calculatePixelColor(surfaceInfo SurfaceHitInfo, camera Camera) Color {
	ambientLightDirection := Normalize(Sub(ambientLightPosition, surfaceInfo.Point))
	viewDirection := Normalize(Sub(camera.Position, surfaceInfo.Point))
	halfwayDirection := Normalize(Add(ambientLightDirection, viewDirection))

	ambient := MultiplyScalar(ambientColor.AsNormalizedVec3(), AmbientStrength)
	diffuse := MultiplyScalar(ambientColor.AsNormalizedVec3(), math.Max(Dot(surfaceInfo.Normal, ambientLightDirection), 0.0))
	specular := MultiplyScalar(ambientColor.AsNormalizedVec3(), math.Pow(math.Max(Dot(surfaceInfo.Normal, halfwayDirection), 0.0), SpecularShininess)*SpecularStrength)

	lightColor := AddAll(ambient, diffuse, specular)
	outputColor := Multiply(lightColor, surfaceColor.AsNormalizedVec3())

	return ColorFromNormalizedVec3(outputColor)
}

// Worker GoRoutine that fetches a render task from the queue and executes it
func renderWorker(waitGroup *sync.WaitGroup, pendingWorkQueue <-chan RenderTask, finishedWorkQueue chan<- RenderResult, id int) {
	defer waitGroup.Done()

	for {
		// Block until either a task is found, or until the channel has been closed
		task, more := <-pendingWorkQueue

		if more {
			log.Println("Worker", id, "started on a task from the pending work queue - remaining tasks:", len(pendingWorkQueue))
			finishedWorkQueue <- render(task)
		} else {
			log.Println("Worker", id, "ran out of tasks - shutting down GoRoutine now")
			break
		}
	}
}

// Render the scene
func render(task RenderTask) RenderResult {
	renderResult := RenderResult{task.StartRow, task.RowCount, make([]Color, task.RowCount*ImageResolutionX)}

	pixelIndex := 0
	for y := task.StartRow; y < task.StartRow+task.RowCount; y++ {
		for x := 0; x < ImageResolutionX; x++ {
			ray := camera.GenerateRayForPixelCenter(x, y, ImageResolutionX, ImageResolutionY)
			pixelColor := Color{Red: 0, Green: 0, Blue: 0, Alpha: 0}

			didHit, hitInfo := camera.MarchAlongRay(ray, scene, RayStepSize)

			if didHit {
				pixelColor = calculatePixelColor(hitInfo, camera)
			}

			renderResult.Pixels[pixelIndex] = pixelColor
			pixelIndex++
		}
	}

	return renderResult
}

// Wait for all worker GoRoutines to finish and close the finished work queue afterwards
func waitForWorkers(waitGroup *sync.WaitGroup, finishedWorkQueue chan<- RenderResult) {
	waitGroup.Wait()
	close(finishedWorkQueue)
}

// Application entry point
func main() {
	defer trackTime(time.Now(), "Render")

	image := NewPngImage(ImageResolutionX, ImageResolutionY, ImageFileName)

	// Start all workers
	waitGroup := sync.WaitGroup{}
	workerCount := (runtime.NumCPU() - 1) * GoRoutinesPerAvailableCpu
	totalNumberOfTasks := AverageNumberOfTasksPerWorker * workerCount

	pendingWorkQueue := make(chan RenderTask, totalNumberOfTasks)
	finishedWorkQueue := make(chan RenderResult, totalNumberOfTasks)

	log.Println("Found", runtime.NumCPU(), "CPUs, this will result in", workerCount, "worker GoRoutines")
	log.Println("Each worker has an average of", AverageNumberOfTasksPerWorker, "tasks")
	log.Println("This results in a total of", totalNumberOfTasks, "tasks for all worker GoRoutines combined")
	log.Println("Workload will be divided into", totalNumberOfTasks, "smaller render tasks")
	log.Println("Output image will have a final resolution of", ImageResolutionX, "x", ImageResolutionY, "pixels")

	for i := 0; i < workerCount; i++ {
		waitGroup.Add(1)
		go renderWorker(&waitGroup, pendingWorkQueue, finishedWorkQueue, i)
	}

	// Generate all jobs
	for i := 0; i < totalNumberOfTasks; i++ {
		rowsToRenderPerJob := int(math.Floor(float64(ImageResolutionY) / float64(totalNumberOfTasks)))
		startRow := i * rowsToRenderPerJob
		rowCount := int(math.Max(float64(rowsToRenderPerJob), float64(ImageResolutionY-startRow)))

		pendingWorkQueue <- RenderTask{startRow, rowCount}
	}

	// No more jobs will be added at this point
	close(pendingWorkQueue)

	// Wait until the scene has been rendered
	waitForWorkers(&waitGroup, finishedWorkQueue)

	// Keep reading render results from the queue as the worker GoRoutines slowly finish their work
	for result := range finishedWorkQueue {
		pixelIndex := 0

		// Save each render chunk in the final output file
		for y := result.StartRow; y < result.StartRow+result.RowCount; y++ {
			for x := 0; x < ImageResolutionX; x++ {
				image.SetPixelColor(x, y, result.Pixels[pixelIndex])
				pixelIndex++
			}
		}
	}

	image.WritePngToFile()
}
