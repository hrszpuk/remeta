package test

type Vector3 struct {
	x int
	y int
	z int
}

func Vector3Add(v1, v2 Vector3) Vector3 {
	return Vector3{
		x: v1.x + v2.x,
		y: v1.y + v2.y,
		z: v1.z + v2.z,
	}
}

func (v1 Vector3) Add(v2 Vector3) Vector3 {
	return Vector3Add(v1, v2)
}
