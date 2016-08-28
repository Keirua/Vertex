scene.md

The details of the original image have been lost, but are the information about the scene at some point

var checkboardTexture16 = math3d.CheckboardTexture{16,16}

var whiteMaterial = math3d.Material{math3d.Color01{0.8, 0.8, 0.8}, 0.5, nil}
var redMaterial = math3d.Material{math3d.Color01{1, 0, 0}, 0.5, nil}
var blueMaterial = math3d.Material{math3d.Color01{0, .5, 1}, 0.5, nil}
var purpleMaterial = math3d.Material{math3d.Color01{0.65, .2, 0.97}, 0.5, &checkboardTexture16}
var greenMaterial = math3d.Material{math3d.Color01{0.3, 0.9, 0.2}, 0.5, nil}

var sphereFloor = math3d.Sphere{math3d.Vertex{0, 10003, -20}, 10000.0, whiteMaterial}
var sphere1 = math3d.Sphere{math3d.Vertex{4.0, -1, -5}, 2.0, redMaterial}
var sphere2 = math3d.Sphere{math3d.Vertex{0.0, 0, -15}, 4, greenMaterial}
var sphere3 = math3d.Sphere{math3d.Vertex{3.0, 2, -10}, 3, blueMaterial}
var sphere4 = math3d.Sphere{math3d.Vertex{-5.5, 0, -8}, 3, purpleMaterial}

var light = math3d.Light{math3d.Vertex{3.0, -10, -10}, math3d.Color01{0.65, .6, 0.97}}
var light2 = math3d.Light{math3d.Vertex{0, -5, 0}, math3d.Color01{0.87, 0.8, 0.97}}