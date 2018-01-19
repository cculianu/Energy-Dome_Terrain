package geometry

import (
    "github.com/go-gl/gl/v4.5-core/gl"
    "github.com/go-gl/mathgl/mgl32"
    "unsafe"
    //"math"
    //"fmt"
)

type Geometry struct {
    // Important geometry attributes
    ArrayBuffer     uint32
    VertexBuffer    uint32
    IndexBuffer     uint32
    IndexCount      int32
}

type Object struct {
    Geo     Geometry
    Pos     mgl32.Vec3
    Scale   mgl32.Vec3
    Color   mgl32.Vec3
    IsLight bool
}

type Mesh struct {
    Pos     mgl32.Vec3
    Normal  mgl32.Vec3
    UV      mgl32.Vec2
}

//func setRenderingAttributes(vertexArrayObject, arrayBuffer, location uint32, size int32, normalized bool, stride int32, offset int) {
//    // Find the last bindings so we don't overwrite them
//    var previousVertexArrayObject int32
//    gl.GetIntegerv(gl.VERTEX_ARRAY_BINDING, &previousVertexArrayObject)
//    var previousArrayBuffer int32
//    gl.GetIntegerv(gl.ARRAY_BUFFER, &previousArrayBuffer)
//
//    // Set our vertex attributes and pointers
//    gl.BindVertexArray(vertexArrayObject)
//    gl.BindBuffer(gl.ARRAY_BUFFER, arrayBuffer)
//    gl.EnableVertexAttribArray(location)
//    gl.VertexAttribPointer(location, size, gl.FLOAT, normalized, stride, gl.PtrOffset(offset))
//
//    // Reset the old bindings.
//    gl.BindBuffer(gl.ARRAY_BUFFER, uint32(previousArrayBuffer))
//    gl.BindVertexArray(uint32(previousVertexArrayObject))
//}
//
//func createRectangle(vertices []float32) Geometry {
//    var rectangle = Geometry{}
//
//    rectangle.VertexCount = 6
//
//    gl.GenBuffers(1, &rectangle.ArrayBuffer)
//    gl.BindBuffer(gl.ARRAY_BUFFER, rectangle.ArrayBuffer)
//    gl.BufferData(gl.ARRAY_BUFFER, len(vertices)*4, gl.Ptr(vertices), gl.STATIC_DRAW)
//    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
//
//    gl.GenVertexArrays(1, &rectangle.VertexObject)
//    gl.BindVertexArray(rectangle.VertexObject)
//
//    setRenderingAttributes(rectangle.VertexObject, rectangle.ArrayBuffer, 0, 3, false, 3*4, 0)
//
//    return rectangle
//}
//
//func bilinearPosition(v0, v3, edge01, edge32 mgl32.Vec3, u, v float32) mgl32.Vec3 {
//    point01 := v0.Add(edge01.Mul(u))
//    point32 := v3.Add(edge32.Mul(u))
//    diff    := point32.Sub(point01).Mul(v)
//    return point01.Add(diff)
//}
//
//func createSurfaceVertices(numSubdivisions int, v0, v1, v2, v3 mgl32.Vec3, data []mgl32.Vec3, offset int) {
//    edge01 := v1.Sub(v0)
//    edge32 := v2.Sub(v3)
//    xStep  := 1.0 / float32(numSubdivisions)
//    yStep  := 1.0 / float32(numSubdivisions)
//    quadIndex := 0
//
//    for x := 0; x < numSubdivisions; x+=1 {
//        for y := 0; y < numSubdivisions; y+=1 {
//
//            // bilinear interpolation of the subdivided vertices.
//            p0 := bilinearPosition(v0, v3, edge01, edge32, (float32(x) + 0.0)*xStep, (float32(y) + 0.0)*yStep)
//            p1 := bilinearPosition(v0, v3, edge01, edge32, (float32(x) + 1.0)*xStep, (float32(y) + 0.0)*yStep)
//            p2 := bilinearPosition(v0, v3, edge01, edge32, (float32(x) + 1.0)*xStep, (float32(y) + 1.0)*yStep)
//            p3 := bilinearPosition(v0, v3, edge01, edge32, (float32(x) + 0.0)*xStep, (float32(y) + 1.0)*yStep)
//
//            // 2* because a vertex consists of a position and a normal.
//            data[offset + 2*(quadIndex + 0)] = p0
//            data[offset + 2*(quadIndex + 1)] = p1
//            data[offset + 2*(quadIndex + 2)] = p2
//
//            data[offset + 2*(quadIndex + 3)] = p0
//            data[offset + 2*(quadIndex + 4)] = p2
//            data[offset + 2*(quadIndex + 5)] = p3
//
//            sp10 := p0.Sub(p1)
//            sp12 := p2.Sub(p1)
//            sp30 := p0.Sub(p3)
//            sp32 := p2.Sub(p3)
//
//            // Normal is at position +1 for a given vertex!
//            n1 := sp12.Cross(sp10)
//            data[offset + 2*(quadIndex + 0) + 1] = n1
//            data[offset + 2*(quadIndex + 1) + 1] = n1
//            data[offset + 2*(quadIndex + 2) + 1] = n1
//
//            n2 := sp30.Cross(sp32)
//            data[offset + 2*(quadIndex + 3) + 1] = n2
//            data[offset + 2*(quadIndex + 4) + 1] = n2
//            data[offset + 2*(quadIndex + 5) + 1] = n2
//
//            quadIndex += 6
//        }
//    }
//}
//
//func createUnitCubeVertices(numSubdivisions int, data []mgl32.Vec3) {
//    verticesPerSide := numSubdivisions*numSubdivisions*6
//    vec3dsPerSide   := 2*verticesPerSide
//    createSurfaceVertices(numSubdivisions, mgl32.Vec3{-1, -1,  1}, mgl32.Vec3{ 1, -1,  1}, mgl32.Vec3{ 1,  1,  1}, mgl32.Vec3{-1,  1,  1}, data, 0*vec3dsPerSide) // Front
//    createSurfaceVertices(numSubdivisions, mgl32.Vec3{-1, -1, -1}, mgl32.Vec3{-1, -1,  1}, mgl32.Vec3{-1,  1,  1}, mgl32.Vec3{-1,  1, -1}, data, 1*vec3dsPerSide) // Left
//    createSurfaceVertices(numSubdivisions, mgl32.Vec3{ 1, -1,  1}, mgl32.Vec3{ 1, -1, -1}, mgl32.Vec3{ 1,  1, -1}, mgl32.Vec3{ 1,  1,  1}, data, 2*vec3dsPerSide) // Right
//    createSurfaceVertices(numSubdivisions, mgl32.Vec3{-1,  1,  1}, mgl32.Vec3{ 1,  1,  1}, mgl32.Vec3{ 1,  1, -1}, mgl32.Vec3{-1,  1, -1}, data, 3*vec3dsPerSide) // Top
//    createSurfaceVertices(numSubdivisions, mgl32.Vec3{-1, -1, -1}, mgl32.Vec3{ 1, -1, -1}, mgl32.Vec3{ 1, -1,  1}, mgl32.Vec3{-1, -1,  1}, data, 4*vec3dsPerSide) // Bottom
//    createSurfaceVertices(numSubdivisions, mgl32.Vec3{-1,  1, -1}, mgl32.Vec3{ 1,  1, -1}, mgl32.Vec3{ 1, -1, -1}, mgl32.Vec3{-1, -1, -1}, data, 5*vec3dsPerSide) // Back
//}
//
//func CreateSurface(numSubdivisions int) Geometry {
//    v0, v1, v2, v3 := mgl32.Vec3{-0.5,0,-0.5}, mgl32.Vec3{0.5,0,-0.5}, mgl32.Vec3{0.5,0,0.5}, mgl32.Vec3{-0.5,0,0.5}
//    numVertices := int32(6*numSubdivisions*numSubdivisions)
//
//    emptyVec := mgl32.Vec3{}
//    stride := int(unsafe.Sizeof(emptyVec))
//    byteSizeVertex := int32(2*stride)
//    byteSizeData   := numVertices*byteSizeVertex
//    data := make([]mgl32.Vec3, byteSizeData, byteSizeData)
//
//    createSurfaceVertices(numSubdivisions, v0, v1, v2, v3, data, 0)
//
//    geometry := Geometry{}
//    geometry.VertexCount = numVertices
//
//    gl.GenBuffers(1, &geometry.ArrayBuffer)
//    gl.BindBuffer(gl.ARRAY_BUFFER, geometry.ArrayBuffer)
//    gl.BufferData(gl.ARRAY_BUFFER, int(byteSizeData), gl.Ptr(data), gl.STATIC_DRAW)
//    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
//
//    gl.GenVertexArrays(1, &geometry.VertexObject)
//    gl.BindVertexArray(geometry.VertexObject)
//
//    setRenderingAttributes(geometry.VertexObject, geometry.ArrayBuffer, 0, 3, false, byteSizeVertex, 0)
//    setRenderingAttributes(geometry.VertexObject, geometry.ArrayBuffer, 1, 3, true , byteSizeVertex, stride)
//
//
//    return geometry
//}
//
//func CreateUnitCube(numSubdivisions int) Geometry {
//    numVerticesPerSide := 6*numSubdivisions*numSubdivisions
//    numVertices := int32(6*numVerticesPerSide)
//
//    emptyVec := mgl32.Vec3{}
//    stride := int(unsafe.Sizeof(emptyVec))
//    byteSizeVertex := int32(stride)
//    // 2 because of there are 2x vec3 (pos+normal).
//    byteSizeData := numVertices*byteSizeVertex*2
//    data := make([]mgl32.Vec3, byteSizeData, byteSizeData)
//
//    createUnitCubeVertices(numSubdivisions, data)
//
//    geometry := Geometry{}
//    geometry.VertexCount = numVertices
//
//    gl.GenBuffers(1, &geometry.ArrayBuffer)
//    gl.BindBuffer(gl.ARRAY_BUFFER, geometry.ArrayBuffer)
//    gl.BufferData(gl.ARRAY_BUFFER, int(byteSizeData), gl.Ptr(data), gl.STATIC_DRAW)
//    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
//
//    gl.GenVertexArrays(1, &geometry.VertexObject)
//    gl.BindVertexArray(geometry.VertexObject)
//
//    // 2 because of there are 2x vec3 (pos+normal).
//    setRenderingAttributes(geometry.VertexObject, geometry.ArrayBuffer, 0, 3, false, byteSizeVertex*2, 0)
//    setRenderingAttributes(geometry.VertexObject, geometry.ArrayBuffer, 1, 3, true , byteSizeVertex*2, stride)
//
//    return geometry
//}
//
//func CreateUnitSphere(numSubdivisions int) Geometry {
//    numVerticesPerSide := 6*numSubdivisions*numSubdivisions
//    numVertices := int32(6*numVerticesPerSide)
//
//    emptyVec := mgl32.Vec3{}
//    stride := int(unsafe.Sizeof(emptyVec))
//    byteSizeVertex := int32(2*stride)
//    byteSizeData := numVertices*byteSizeVertex
//
//    data := make([]mgl32.Vec3, byteSizeData, byteSizeData)
//
//    createUnitCubeVertices(numSubdivisions, data)
//
//    for i := 0; i < int(numVertices); i+=1 {
//        // Cubical position
//        c := data[2*i]
//
//        // Spherical position
//        s := mgl32.Vec3{}
//        s[0] = c.X() * float32(math.Sqrt(float64(1.0 - (c.Y() * c.Y()) / 2.0 - (c.Z() * c.Z()) / 2.0 + (c.Y() * c.Y() * c.Z() * c.Z()) / 3.0)))
//        s[1] = c.Y() * float32(math.Sqrt(float64(1.0 - (c.Z() * c.Z()) / 2.0 - (c.X() * c.X()) / 2.0 + (c.Z() * c.Z() * c.X() * c.X()) / 3.0)))
//        s[2] = c.Z() * float32(math.Sqrt(float64(1.0 - (c.X() * c.X()) / 2.0 - (c.Y() * c.Y()) / 2.0 + (c.X() * c.X() * c.Y() * c.Y()) / 3.0)))
//
//        data[2*i] = s
//
//        // For a unit sphere, the normal is equal the position!
//        data[2*i + 1] = s
//    }
//
//    geometry := Geometry{}
//    geometry.VertexCount = numVertices
//
//    gl.GenBuffers(1, &geometry.ArrayBuffer)
//    gl.BindBuffer(gl.ARRAY_BUFFER, geometry.ArrayBuffer)
//    gl.BufferData(gl.ARRAY_BUFFER, int(byteSizeData), gl.Ptr(data), gl.STATIC_DRAW)
//    gl.BindBuffer(gl.ARRAY_BUFFER, 0)
//
//    gl.GenVertexArrays(1, &geometry.VertexObject)
//    gl.BindVertexArray(geometry.VertexObject)
//
//    setRenderingAttributes(geometry.VertexObject, geometry.ArrayBuffer, 0, 3, false, byteSizeVertex, 0)
//    setRenderingAttributes(geometry.VertexObject, geometry.ArrayBuffer, 1, 3, true , byteSizeVertex, stride)
//
//    return geometry
//}

func CreateUnitSquare(subdivisions int, offset mgl32.Vec2) Geometry {
    size := subdivisions+2
    vertexCount := size*size

    mesh := make([]Mesh, vertexCount, vertexCount)

    sizeFactor := 1.0/float32(size-1)

    for i := 0; i < size; i++ {
        for j := 0; j < size; j++ {
            // It's now centered at the origin on both axes and has unit size.
            // Possibly with an offset defined.
            mesh[j*size+i].Pos    = mgl32.Vec3{float32(i) * sizeFactor - 0.5 + offset.X(), 0, float32(j) * sizeFactor - 0.5 + offset.Y()}
            mesh[j*size+i].Normal = mgl32.Vec3{0, 1, 0}
            mesh[j*size+i].UV     = mgl32.Vec2{float32(i)*sizeFactor, float32(j)*sizeFactor}
        }
    }

    indexCount  := (size-1)*(size-1)*2*3
    meshIndices := make([]uint32, indexCount, indexCount)

    index := 0
    uSize := uint32(size)
    for i := uint32(0); i < uSize-1; i++ {
        for j := uint32(0); j < uSize-1; j++ {
            meshIndices[index]   = j*uSize+i
            meshIndices[index+1] = j*uSize+i+1
            meshIndices[index+2] = (j+1)*uSize+i+1

            meshIndices[index+3] = j*uSize+i
            meshIndices[index+4] = (j+1)*uSize+i+1
            meshIndices[index+5] = (j+1)*uSize+i
            index += 6
        }
    }

    geo := Geometry{}

    var m Mesh
    stride := int32(unsafe.Sizeof(m))
    var v mgl32.Vec3
    vStride := int(unsafe.Sizeof(v))
    var ui uint32
    uiStride := int(unsafe.Sizeof(ui))

    gl.GenBuffers(1, &geo.ArrayBuffer)
    gl.BindBuffer(gl.ARRAY_BUFFER, geo.ArrayBuffer)
    //gl.BufferData(gl.ARRAY_BUFFER, int(stride)*vertexCount, gl.Ptr(&mesh[0].Pos[0]), gl.STATIC_DRAW)
    gl.BufferData(gl.ARRAY_BUFFER, int(stride)*vertexCount, gl.Ptr(mesh), gl.STATIC_DRAW)

    gl.GenVertexArrays(1, &geo.VertexBuffer)
    gl.BindVertexArray(geo.VertexBuffer)

    gl.EnableVertexAttribArray(0)
    gl.VertexAttribPointer(0, 3, gl.FLOAT, false, stride, gl.PtrOffset(0))
    gl.EnableVertexAttribArray(1)
    gl.VertexAttribPointer(1, 3, gl.FLOAT, true, stride, gl.PtrOffset(vStride))
    gl.EnableVertexAttribArray(2)
    gl.VertexAttribPointer(2, 2, gl.FLOAT, false, stride, gl.PtrOffset(2*vStride))

    gl.BindBuffer (gl.ARRAY_BUFFER, 0)

    gl.GenBuffers (1, &geo.IndexBuffer)
    gl.BindBuffer (gl.ELEMENT_ARRAY_BUFFER, geo.IndexBuffer)
    gl.BufferData (gl.ELEMENT_ARRAY_BUFFER, uiStride*indexCount, gl.Ptr(meshIndices), gl.STATIC_DRAW)
    gl.BindBuffer (gl.ELEMENT_ARRAY_BUFFER, 0)

    geo.IndexCount = int32(indexCount)

    return geo;
}


func CreateObject(geo Geometry, pos, scale, color mgl32.Vec3, isLight bool) Object {
    return Object {
        Geo: geo,
        Pos: pos,
        Scale: scale.Mul(0.5),
        Color: color,
        IsLight: isLight,
    }
}

