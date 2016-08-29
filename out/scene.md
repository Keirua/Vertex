scene.md

The details of the original image have been lost, but are the information about the scene at some point :

var checkboardTexture16 = CheckboardTexture{16,16}
var imageTexture = FileTexture{"images/cropped-P1120606-small.jpg",nil}

var turbulence = NewTurbulence(32)
var marble5 = NewMarble(5, 10, 5, 32)
var marble1 = NewMarble(5, 10, 1, 32)
var wood = NewWood(12, 0.1, 32)

var whiteMaterial = Material{Color01{0.8, 0.8, 0.8}, 0.5, nil}
var redMaterial = Material{Color01{1, 0, 0}, 0.5, nil}
var blueMaterial = Material{Color01{0, .5, 1}, 0.5, marble5}
var purpleMaterial = Material{Color01{0.65, .2, 0.97}, 0.5, &checkboardTexture16}
var greenMaterial = Material{Color01{0.3, 0.9, 0.2}, 0.5, turbulence}

var sphereFloor = Sphere{Vertex{0, 10003, -20}, 10000.0, whiteMaterial}
var sphere1 = Sphere{Vertex{4.0, -1, -5}, 2.0, redMaterial}
var sphere2 = Sphere{Vertex{0.0, 0, -15}, 4, greenMaterial}
var sphere3 = Sphere{Vertex{3.0, 2, -10}, 3, blueMaterial}
var sphere4 = Sphere{Vertex{-5.5, 0, -8}, 3, purpleMaterial}

var light = Light{Vertex{3.0, -10, -10}, Color01{0.65, .6, 0.97}}
var light2 = Light{Vertex{0, -5, 0}, Color01{0.87, 0.8, 0.97}}