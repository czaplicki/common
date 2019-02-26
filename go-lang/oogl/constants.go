package oogl

import	"github.com/go-gl/gl/v4.1-core/gl"

const ( //Types
	INT			= gl.INT
	UINT		= gl.UNSIGNED_INT
	FLOAT		= gl.FLOAT
)


const ( //Clear masks
	COLOR_BUFFER_BIT = gl.COLOR_BUFFER_BIT
	DEPTH_BUFFER_BIT = gl.DEPTH_BUFFER_BIT
)

const ( //drawModes
	TRIANGLES = gl.TRIANGLES
)
const ( //Texture slots
	TEXTURE0				= gl.TEXTURE0
	TEXTURE1				= gl.TEXTURE1
	TEXTURE2				= gl.TEXTURE2
	TEXTURE3				= gl.TEXTURE3
)
const ( //Texture Options
	TEXTURE_MIN_FILTER        			= gl.TEXTURE_MIN_FILTER
	TEXTURE_MAG_FILTER        			= gl.TEXTURE_MAG_FILTER

	TEXTURE_WRAP_S						= gl.TEXTURE_WRAP_S
	TEXTURE_WRAP_T						= gl.TEXTURE_WRAP_T

	//TEXTURE_WRAP_R						= gl.TEXTURE_WRAP_R
	//DEPTH_STENCIL_TEXTURE_MODE			= gl.DEPTH_STENCIL_TEXTURE_MODE
	//TEXTURE_BASE_LEVEL					= gl.TEXTURE_BASE_LEVEL
	//TEXTURE_BORDER_COLOR				= gl.TEXTURE_BORDER_COLOR
	//TEXTURE_COMPARE_FUNC				= gl.TEXTURE_COMPARE_FUNC
	//TEXTURE_COMPARE_MODE				= gl.TEXTURE_COMPARE_MODE
	//TEXTURE_LOD_BIAS					= gl.TEXTURE_LOD_BIAS
	//TEXTURE_MIN_LOD						= gl.TEXTURE_MIN_LOD
	//TEXTURE_MAX_LOD						= gl.TEXTURE_MAX_LOD
	//TEXTURE_MAX_LEVEL					= gl.TEXTURE_MAX_LEVEL
	//TEXTURE_SWIZZLE_R					= gl.TEXTURE_SWIZZLE_R
	//TEXTURE_SWIZZLE_G					= gl.TEXTURE_SWIZZLE_G
	//TEXTURE_SWIZZLE_B					= gl.TEXTURE_SWIZZLE_B
	//TEXTURE_SWIZZLE_A					= gl.TEXTURE_SWIZZLE_A
	//TEXTURE_SWIZZLE_RGBA				= gl.TEXTURE_SWIZZLE_RGBA
)

const (
	//Min and Mag filters
	LINEAR								= gl.LINEAR
	NEAREST								= gl.NEAREST
	//MinFilters
	LINEAR_MIPMAP_LINEAR				= gl.LINEAR_MIPMAP_LINEAR
	LINEAR_MIPMAP_NEAREST				= gl.LINEAR_MIPMAP_NEAREST
	NEAREST_MIPMAP_LINEAR				= gl.NEAREST_MIPMAP_LINEAR
	NEAREST_MIPMAP_NEAREST				= gl.NEAREST_MIPMAP_NEAREST
)

const ( //Wrap S and T
	CLAMP_TO_EDGE						= gl.CLAMP_TO_EDGE
	CLAMP_TO_BORDER						= gl.CLAMP_TO_BORDER
	REPEAT								= gl.REPEAT
	MIRRORED_REPEAT						= gl.MIRRORED_REPEAT
	MIRROR_CLAMP_TO_EDGE				= gl.MIRROR_CLAMP_TO_EDGE
)
