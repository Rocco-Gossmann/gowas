package bmps

import "github.com/rocco-gossmann/GoWas/canvas"

var mem_CursorBMP = []uint32 {
	0x01e06f8b, 0x01e06f8b, 0x01e06f8b, 0x01e06f8b, 0x01e06f8b, 0x01e06f8b, 0x01e06f8b, 0x00000000, 
	0x014d2326, 0x01be2633, 0x01be2633, 0x01be2633, 0x01be2633, 0x01be2633, 0x00000000, 0x00000000, 
	0x014d2326, 0x01be2633, 0x01be2633, 0x01be2633, 0x014d2326, 0x00000000, 0x00000000, 0x00000000, 
	0x014d2326, 0x01be2633, 0x01be2633, 0x01be2633, 0x01be2633, 0x014d2326, 0x00000000, 0x00000000, 
	0x014d2326, 0x01be2633, 0x014d2326, 0x01be2633, 0x01be2633, 0x01be2633, 0x014d2326, 0x00000000, 
	0x014d2326, 0x01be2633, 0x00000000, 0x014d2326, 0x01be2633, 0x01be2633, 0x01be2633, 0x014d2326, 
	0x014d2326, 0x00000000, 0x00000000, 0x00000000, 0x014d2326, 0x01be2633, 0x01be2633, 0x01be2633, 
	0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x00000000, 0x014d2326, 0x01be2633, 0x00000000, 
}

var CursorBMP = canvas.CreateBitmap(8, &mem_CursorBMP)