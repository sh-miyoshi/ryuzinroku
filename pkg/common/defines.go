package common

const (
	// ScreenX is x size of screen
	ScreenX = 640
	// ScreenY is y size of screen
	ScreenY = 480

	// FieldTopX ...
	FieldTopX = 32
	// FiledTopY ...
	FiledTopY = 16
	// FiledSizeX ...
	FiledSizeX = 384
	// FiledSizeY ...
	FiledSizeY = 448
)

// Direct ...
type Direct string

const (
	// DirectFront ...
	DirectFront Direct = "front"
	// DirectLeft ...
	DirectLeft Direct = "left"
	// DirectRight ...
	DirectRight Direct = "right"
)
