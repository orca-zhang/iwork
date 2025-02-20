// Code generated by protoc-gen-go.
// source: TSCH3DArchives.proto
// DO NOT EDIT!

/*
Package TSCH is a generated protocol buffer package.

It is generated from these files:
	TSCH3DArchives.proto

It has these top-level messages:
	Chart3DEnvironmentPackageArchive
	Chart3DFillArchive
	Chart3DPointLightArchive
	Chart3DDirectionalLightArchive
	Chart3DSpotLightArchive
	Chart3DLightArchive
	Chart3DLightingModelArchive
	Chart3DLightingPackageArchive
	Chart3DTexturesMaterialArchive
	Chart3DEmissiveMaterialArchive
	Chart3DDiffuseMaterialArchive
	Chart3DModulateMaterialArchive
	Chart3DSpecularMaterialArchive
	Chart3DShininessMaterialArchive
	Chart3DEnvironmentMaterialArchive
	Chart3DFixedFunctionLightingModelArchive
	Chart3DPhongLightingModelArchive
	Chart3DPhongMaterialPackageArchive
	Chart3DTSPImageDataTextureArchive
	Chart3DBaseImageTextureTilingArchive
	Chart3DImageTextureTilingArchive
	Chart3DVectorArchive
*/
package TSCH

import proto "github.com/golang/protobuf/proto"
import math "math"
import "github.com/orcastor/iwork-converter/proto/TSP"
import "github.com/orcastor/iwork-converter/proto/TSD"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type TextureTilingMode int32

const (
	TextureTilingMode_textureTilingModeNone    TextureTilingMode = 0
	TextureTilingMode_textureTilingModeTallest TextureTilingMode = 1
)

var TextureTilingMode_name = map[int32]string{
	0: "textureTilingModeNone",
	1: "textureTilingModeTallest",
}
var TextureTilingMode_value = map[string]int32{
	"textureTilingModeNone":    0,
	"textureTilingModeTallest": 1,
}

func (x TextureTilingMode) Enum() *TextureTilingMode {
	p := new(TextureTilingMode)
	*p = x
	return p
}
func (x TextureTilingMode) String() string {
	return proto.EnumName(TextureTilingMode_name, int32(x))
}
func (x *TextureTilingMode) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TextureTilingMode_value, data, "TextureTilingMode")
	if err != nil {
		return err
	}
	*x = TextureTilingMode(value)
	return nil
}

type TextureTilingFace int32

const (
	TextureTilingFace_textureTilingFaceAll          TextureTilingFace = 0
	TextureTilingFace_textureTilingFaceTopAndBottom TextureTilingFace = 1
	TextureTilingFace_textureTilingFaceSide         TextureTilingFace = 2
)

var TextureTilingFace_name = map[int32]string{
	0: "textureTilingFaceAll",
	1: "textureTilingFaceTopAndBottom",
	2: "textureTilingFaceSide",
}
var TextureTilingFace_value = map[string]int32{
	"textureTilingFaceAll":          0,
	"textureTilingFaceTopAndBottom": 1,
	"textureTilingFaceSide":         2,
}

func (x TextureTilingFace) Enum() *TextureTilingFace {
	p := new(TextureTilingFace)
	*p = x
	return p
}
func (x TextureTilingFace) String() string {
	return proto.EnumName(TextureTilingFace_name, int32(x))
}
func (x *TextureTilingFace) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TextureTilingFace_value, data, "TextureTilingFace")
	if err != nil {
		return err
	}
	*x = TextureTilingFace(value)
	return nil
}

type TextureTilingWrap int32

const (
	TextureTilingWrap_textureTilingWrapProjected TextureTilingWrap = 0
	TextureTilingWrap_textureTilingWrapFaceWrap  TextureTilingWrap = 1
)

var TextureTilingWrap_name = map[int32]string{
	0: "textureTilingWrapProjected",
	1: "textureTilingWrapFaceWrap",
}
var TextureTilingWrap_value = map[string]int32{
	"textureTilingWrapProjected": 0,
	"textureTilingWrapFaceWrap":  1,
}

func (x TextureTilingWrap) Enum() *TextureTilingWrap {
	p := new(TextureTilingWrap)
	*p = x
	return p
}
func (x TextureTilingWrap) String() string {
	return proto.EnumName(TextureTilingWrap_name, int32(x))
}
func (x *TextureTilingWrap) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TextureTilingWrap_value, data, "TextureTilingWrap")
	if err != nil {
		return err
	}
	*x = TextureTilingWrap(value)
	return nil
}

type TextureTilingXPosition int32

const (
	TextureTilingXPosition_textureTilingXPositionLeft   TextureTilingXPosition = 0
	TextureTilingXPosition_textureTilingXPositionCenter TextureTilingXPosition = 1
	TextureTilingXPosition_textureTilingXPositionRight  TextureTilingXPosition = 2
)

var TextureTilingXPosition_name = map[int32]string{
	0: "textureTilingXPositionLeft",
	1: "textureTilingXPositionCenter",
	2: "textureTilingXPositionRight",
}
var TextureTilingXPosition_value = map[string]int32{
	"textureTilingXPositionLeft":   0,
	"textureTilingXPositionCenter": 1,
	"textureTilingXPositionRight":  2,
}

func (x TextureTilingXPosition) Enum() *TextureTilingXPosition {
	p := new(TextureTilingXPosition)
	*p = x
	return p
}
func (x TextureTilingXPosition) String() string {
	return proto.EnumName(TextureTilingXPosition_name, int32(x))
}
func (x *TextureTilingXPosition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TextureTilingXPosition_value, data, "TextureTilingXPosition")
	if err != nil {
		return err
	}
	*x = TextureTilingXPosition(value)
	return nil
}

type TextureTilingYPosition int32

const (
	TextureTilingYPosition_textureTilingYPositionTop    TextureTilingYPosition = 0
	TextureTilingYPosition_textureTilingYPositionMiddle TextureTilingYPosition = 1
	TextureTilingYPosition_textureTilingYPositionBottom TextureTilingYPosition = 2
)

var TextureTilingYPosition_name = map[int32]string{
	0: "textureTilingYPositionTop",
	1: "textureTilingYPositionMiddle",
	2: "textureTilingYPositionBottom",
}
var TextureTilingYPosition_value = map[string]int32{
	"textureTilingYPositionTop":    0,
	"textureTilingYPositionMiddle": 1,
	"textureTilingYPositionBottom": 2,
}

func (x TextureTilingYPosition) Enum() *TextureTilingYPosition {
	p := new(TextureTilingYPosition)
	*p = x
	return p
}
func (x TextureTilingYPosition) String() string {
	return proto.EnumName(TextureTilingYPosition_name, int32(x))
}
func (x *TextureTilingYPosition) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TextureTilingYPosition_value, data, "TextureTilingYPosition")
	if err != nil {
		return err
	}
	*x = TextureTilingYPosition(value)
	return nil
}

type TextureTilingContinuity int32

const (
	TextureTilingContinuity_textureTilingContinuityNone     TextureTilingContinuity = 0
	TextureTilingContinuity_textureTilingContinuityGlobal   TextureTilingContinuity = 1
	TextureTilingContinuity_textureTilingContinuitySeries   TextureTilingContinuity = 2
	TextureTilingContinuity_textureTilingContinuityJittered TextureTilingContinuity = 3
)

var TextureTilingContinuity_name = map[int32]string{
	0: "textureTilingContinuityNone",
	1: "textureTilingContinuityGlobal",
	2: "textureTilingContinuitySeries",
	3: "textureTilingContinuityJittered",
}
var TextureTilingContinuity_value = map[string]int32{
	"textureTilingContinuityNone":     0,
	"textureTilingContinuityGlobal":   1,
	"textureTilingContinuitySeries":   2,
	"textureTilingContinuityJittered": 3,
}

func (x TextureTilingContinuity) Enum() *TextureTilingContinuity {
	p := new(TextureTilingContinuity)
	*p = x
	return p
}
func (x TextureTilingContinuity) String() string {
	return proto.EnumName(TextureTilingContinuity_name, int32(x))
}
func (x *TextureTilingContinuity) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(TextureTilingContinuity_value, data, "TextureTilingContinuity")
	if err != nil {
		return err
	}
	*x = TextureTilingContinuity(value)
	return nil
}

type FillPropertyType int32

const (
	FillPropertyType_fillPropertyTypeUndefined FillPropertyType = 0
	FillPropertyType_fillPropertyTypeArea      FillPropertyType = 1
	FillPropertyType_fillPropertyTypeBar       FillPropertyType = 2
	FillPropertyType_fillPropertyTypeColumn    FillPropertyType = 3
	FillPropertyType_fillPropertyTypeLine      FillPropertyType = 4
	FillPropertyType_fillPropertyTypePie       FillPropertyType = 5
)

var FillPropertyType_name = map[int32]string{
	0: "fillPropertyTypeUndefined",
	1: "fillPropertyTypeArea",
	2: "fillPropertyTypeBar",
	3: "fillPropertyTypeColumn",
	4: "fillPropertyTypeLine",
	5: "fillPropertyTypePie",
}
var FillPropertyType_value = map[string]int32{
	"fillPropertyTypeUndefined": 0,
	"fillPropertyTypeArea":      1,
	"fillPropertyTypeBar":       2,
	"fillPropertyTypeColumn":    3,
	"fillPropertyTypeLine":      4,
	"fillPropertyTypePie":       5,
}

func (x FillPropertyType) Enum() *FillPropertyType {
	p := new(FillPropertyType)
	*p = x
	return p
}
func (x FillPropertyType) String() string {
	return proto.EnumName(FillPropertyType_name, int32(x))
}
func (x *FillPropertyType) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(FillPropertyType_value, data, "FillPropertyType")
	if err != nil {
		return err
	}
	*x = FillPropertyType(value)
	return nil
}

type Chart3DEnvironmentPackageArchive struct {
	Materials        []*Chart3DEnvironmentMaterialArchive `protobuf:"bytes,1,rep,name=materials" json:"materials,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *Chart3DEnvironmentPackageArchive) Reset()         { *m = Chart3DEnvironmentPackageArchive{} }
func (m *Chart3DEnvironmentPackageArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DEnvironmentPackageArchive) ProtoMessage()    {}

func (m *Chart3DEnvironmentPackageArchive) GetMaterials() []*Chart3DEnvironmentMaterialArchive {
	if m != nil {
		return m.Materials
	}
	return nil
}

type Chart3DFillArchive struct {
	Lightingmodel    *Chart3DLightingModelArchive `protobuf:"bytes,1,opt,name=lightingmodel" json:"lightingmodel,omitempty"`
	TexturesetId     *string                      `protobuf:"bytes,2,opt,name=textureset_id" json:"textureset_id,omitempty"`
	FillType         *FillPropertyType            `protobuf:"varint,3,opt,name=fill_type,enum=TSCH.FillPropertyType" json:"fill_type,omitempty"`
	SeriesIndex      *uint32                      `protobuf:"varint,4,opt,name=series_index" json:"series_index,omitempty"`
	XXX_unrecognized []byte                       `json:"-"`
}

func (m *Chart3DFillArchive) Reset()         { *m = Chart3DFillArchive{} }
func (m *Chart3DFillArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DFillArchive) ProtoMessage()    {}

func (m *Chart3DFillArchive) GetLightingmodel() *Chart3DLightingModelArchive {
	if m != nil {
		return m.Lightingmodel
	}
	return nil
}

func (m *Chart3DFillArchive) GetTexturesetId() string {
	if m != nil && m.TexturesetId != nil {
		return *m.TexturesetId
	}
	return ""
}

func (m *Chart3DFillArchive) GetFillType() FillPropertyType {
	if m != nil && m.FillType != nil {
		return *m.FillType
	}
	return FillPropertyType_fillPropertyTypeUndefined
}

func (m *Chart3DFillArchive) GetSeriesIndex() uint32 {
	if m != nil && m.SeriesIndex != nil {
		return *m.SeriesIndex
	}
	return 0
}

type Chart3DPointLightArchive struct {
	Position         *Chart3DVectorArchive `protobuf:"bytes,1,req,name=position" json:"position,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Chart3DPointLightArchive) Reset()         { *m = Chart3DPointLightArchive{} }
func (m *Chart3DPointLightArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DPointLightArchive) ProtoMessage()    {}

func (m *Chart3DPointLightArchive) GetPosition() *Chart3DVectorArchive {
	if m != nil {
		return m.Position
	}
	return nil
}

type Chart3DDirectionalLightArchive struct {
	Direction        *Chart3DVectorArchive `protobuf:"bytes,1,req,name=direction" json:"direction,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Chart3DDirectionalLightArchive) Reset()         { *m = Chart3DDirectionalLightArchive{} }
func (m *Chart3DDirectionalLightArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DDirectionalLightArchive) ProtoMessage()    {}

func (m *Chart3DDirectionalLightArchive) GetDirection() *Chart3DVectorArchive {
	if m != nil {
		return m.Direction
	}
	return nil
}

type Chart3DSpotLightArchive struct {
	Position         *Chart3DVectorArchive `protobuf:"bytes,1,req,name=position" json:"position,omitempty"`
	Direction        *Chart3DVectorArchive `protobuf:"bytes,2,req,name=direction" json:"direction,omitempty"`
	Cutoff           *float32              `protobuf:"fixed32,3,req,name=cutoff" json:"cutoff,omitempty"`
	Dropoff          *float32              `protobuf:"fixed32,4,req,name=dropoff" json:"dropoff,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Chart3DSpotLightArchive) Reset()         { *m = Chart3DSpotLightArchive{} }
func (m *Chart3DSpotLightArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DSpotLightArchive) ProtoMessage()    {}

func (m *Chart3DSpotLightArchive) GetPosition() *Chart3DVectorArchive {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Chart3DSpotLightArchive) GetDirection() *Chart3DVectorArchive {
	if m != nil {
		return m.Direction
	}
	return nil
}

func (m *Chart3DSpotLightArchive) GetCutoff() float32 {
	if m != nil && m.Cutoff != nil {
		return *m.Cutoff
	}
	return 0
}

func (m *Chart3DSpotLightArchive) GetDropoff() float32 {
	if m != nil && m.Dropoff != nil {
		return *m.Dropoff
	}
	return 0
}

type Chart3DLightArchive struct {
	Name             *string                         `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	AmbientColor     *Chart3DVectorArchive           `protobuf:"bytes,2,req,name=ambient_color" json:"ambient_color,omitempty"`
	DiffuseColor     *Chart3DVectorArchive           `protobuf:"bytes,3,req,name=diffuse_color" json:"diffuse_color,omitempty"`
	SpecularColor    *Chart3DVectorArchive           `protobuf:"bytes,4,req,name=specular_color" json:"specular_color,omitempty"`
	Intensity        *float32                        `protobuf:"fixed32,5,req,name=intensity" json:"intensity,omitempty"`
	Attenuation      *Chart3DVectorArchive           `protobuf:"bytes,6,req,name=attenuation" json:"attenuation,omitempty"`
	CoordinateSpace  *uint32                         `protobuf:"varint,7,req,name=coordinate_space" json:"coordinate_space,omitempty"`
	Enabled          *bool                           `protobuf:"varint,8,req,name=enabled" json:"enabled,omitempty"`
	PointLight       *Chart3DPointLightArchive       `protobuf:"bytes,9,opt,name=point_light" json:"point_light,omitempty"`
	DirectionalLight *Chart3DDirectionalLightArchive `protobuf:"bytes,10,opt,name=directional_light" json:"directional_light,omitempty"`
	SpotLight        *Chart3DSpotLightArchive        `protobuf:"bytes,11,opt,name=spot_light" json:"spot_light,omitempty"`
	XXX_unrecognized []byte                          `json:"-"`
}

func (m *Chart3DLightArchive) Reset()         { *m = Chart3DLightArchive{} }
func (m *Chart3DLightArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DLightArchive) ProtoMessage()    {}

func (m *Chart3DLightArchive) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Chart3DLightArchive) GetAmbientColor() *Chart3DVectorArchive {
	if m != nil {
		return m.AmbientColor
	}
	return nil
}

func (m *Chart3DLightArchive) GetDiffuseColor() *Chart3DVectorArchive {
	if m != nil {
		return m.DiffuseColor
	}
	return nil
}

func (m *Chart3DLightArchive) GetSpecularColor() *Chart3DVectorArchive {
	if m != nil {
		return m.SpecularColor
	}
	return nil
}

func (m *Chart3DLightArchive) GetIntensity() float32 {
	if m != nil && m.Intensity != nil {
		return *m.Intensity
	}
	return 0
}

func (m *Chart3DLightArchive) GetAttenuation() *Chart3DVectorArchive {
	if m != nil {
		return m.Attenuation
	}
	return nil
}

func (m *Chart3DLightArchive) GetCoordinateSpace() uint32 {
	if m != nil && m.CoordinateSpace != nil {
		return *m.CoordinateSpace
	}
	return 0
}

func (m *Chart3DLightArchive) GetEnabled() bool {
	if m != nil && m.Enabled != nil {
		return *m.Enabled
	}
	return false
}

func (m *Chart3DLightArchive) GetPointLight() *Chart3DPointLightArchive {
	if m != nil {
		return m.PointLight
	}
	return nil
}

func (m *Chart3DLightArchive) GetDirectionalLight() *Chart3DDirectionalLightArchive {
	if m != nil {
		return m.DirectionalLight
	}
	return nil
}

func (m *Chart3DLightArchive) GetSpotLight() *Chart3DSpotLightArchive {
	if m != nil {
		return m.SpotLight
	}
	return nil
}

type Chart3DLightingModelArchive struct {
	Phong            *Chart3DPhongLightingModelArchive         `protobuf:"bytes,1,opt,name=phong" json:"phong,omitempty"`
	FixedFunction    *Chart3DFixedFunctionLightingModelArchive `protobuf:"bytes,2,opt,name=fixed_function" json:"fixed_function,omitempty"`
	Environment      *Chart3DEnvironmentPackageArchive         `protobuf:"bytes,3,opt,name=environment" json:"environment,omitempty"`
	XXX_unrecognized []byte                                    `json:"-"`
}

func (m *Chart3DLightingModelArchive) Reset()         { *m = Chart3DLightingModelArchive{} }
func (m *Chart3DLightingModelArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DLightingModelArchive) ProtoMessage()    {}

func (m *Chart3DLightingModelArchive) GetPhong() *Chart3DPhongLightingModelArchive {
	if m != nil {
		return m.Phong
	}
	return nil
}

func (m *Chart3DLightingModelArchive) GetFixedFunction() *Chart3DFixedFunctionLightingModelArchive {
	if m != nil {
		return m.FixedFunction
	}
	return nil
}

func (m *Chart3DLightingModelArchive) GetEnvironment() *Chart3DEnvironmentPackageArchive {
	if m != nil {
		return m.Environment
	}
	return nil
}

type Chart3DLightingPackageArchive struct {
	Name             *string                `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Lights           []*Chart3DLightArchive `protobuf:"bytes,2,rep,name=lights" json:"lights,omitempty"`
	XXX_unrecognized []byte                 `json:"-"`
}

func (m *Chart3DLightingPackageArchive) Reset()         { *m = Chart3DLightingPackageArchive{} }
func (m *Chart3DLightingPackageArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DLightingPackageArchive) ProtoMessage()    {}

func (m *Chart3DLightingPackageArchive) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *Chart3DLightingPackageArchive) GetLights() []*Chart3DLightArchive {
	if m != nil {
		return m.Lights
	}
	return nil
}

type Chart3DTexturesMaterialArchive struct {
	Color            *Chart3DVectorArchive                `protobuf:"bytes,1,req,name=color" json:"color,omitempty"`
	Textures         []*Chart3DTSPImageDataTextureArchive `protobuf:"bytes,2,rep,name=textures" json:"textures,omitempty"`
	XXX_unrecognized []byte                               `json:"-"`
}

func (m *Chart3DTexturesMaterialArchive) Reset()         { *m = Chart3DTexturesMaterialArchive{} }
func (m *Chart3DTexturesMaterialArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DTexturesMaterialArchive) ProtoMessage()    {}

func (m *Chart3DTexturesMaterialArchive) GetColor() *Chart3DVectorArchive {
	if m != nil {
		return m.Color
	}
	return nil
}

func (m *Chart3DTexturesMaterialArchive) GetTextures() []*Chart3DTSPImageDataTextureArchive {
	if m != nil {
		return m.Textures
	}
	return nil
}

type Chart3DEmissiveMaterialArchive struct {
	Super            *Chart3DTexturesMaterialArchive     `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	Tilings          []*Chart3DImageTextureTilingArchive `protobuf:"bytes,2,rep,name=tilings" json:"tilings,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *Chart3DEmissiveMaterialArchive) Reset()         { *m = Chart3DEmissiveMaterialArchive{} }
func (m *Chart3DEmissiveMaterialArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DEmissiveMaterialArchive) ProtoMessage()    {}

func (m *Chart3DEmissiveMaterialArchive) GetSuper() *Chart3DTexturesMaterialArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *Chart3DEmissiveMaterialArchive) GetTilings() []*Chart3DImageTextureTilingArchive {
	if m != nil {
		return m.Tilings
	}
	return nil
}

type Chart3DDiffuseMaterialArchive struct {
	Super            *Chart3DTexturesMaterialArchive     `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	Tilings          []*Chart3DImageTextureTilingArchive `protobuf:"bytes,2,rep,name=tilings" json:"tilings,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *Chart3DDiffuseMaterialArchive) Reset()         { *m = Chart3DDiffuseMaterialArchive{} }
func (m *Chart3DDiffuseMaterialArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DDiffuseMaterialArchive) ProtoMessage()    {}

func (m *Chart3DDiffuseMaterialArchive) GetSuper() *Chart3DTexturesMaterialArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *Chart3DDiffuseMaterialArchive) GetTilings() []*Chart3DImageTextureTilingArchive {
	if m != nil {
		return m.Tilings
	}
	return nil
}

type Chart3DModulateMaterialArchive struct {
	Super            *Chart3DTexturesMaterialArchive     `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	Tilings          []*Chart3DImageTextureTilingArchive `protobuf:"bytes,2,rep,name=tilings" json:"tilings,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *Chart3DModulateMaterialArchive) Reset()         { *m = Chart3DModulateMaterialArchive{} }
func (m *Chart3DModulateMaterialArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DModulateMaterialArchive) ProtoMessage()    {}

func (m *Chart3DModulateMaterialArchive) GetSuper() *Chart3DTexturesMaterialArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *Chart3DModulateMaterialArchive) GetTilings() []*Chart3DImageTextureTilingArchive {
	if m != nil {
		return m.Tilings
	}
	return nil
}

type Chart3DSpecularMaterialArchive struct {
	Super            *Chart3DTexturesMaterialArchive     `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	Tilings          []*Chart3DImageTextureTilingArchive `protobuf:"bytes,2,rep,name=tilings" json:"tilings,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *Chart3DSpecularMaterialArchive) Reset()         { *m = Chart3DSpecularMaterialArchive{} }
func (m *Chart3DSpecularMaterialArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DSpecularMaterialArchive) ProtoMessage()    {}

func (m *Chart3DSpecularMaterialArchive) GetSuper() *Chart3DTexturesMaterialArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *Chart3DSpecularMaterialArchive) GetTilings() []*Chart3DImageTextureTilingArchive {
	if m != nil {
		return m.Tilings
	}
	return nil
}

type Chart3DShininessMaterialArchive struct {
	Super            *Chart3DTexturesMaterialArchive     `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	Tilings          []*Chart3DImageTextureTilingArchive `protobuf:"bytes,2,rep,name=tilings" json:"tilings,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *Chart3DShininessMaterialArchive) Reset()         { *m = Chart3DShininessMaterialArchive{} }
func (m *Chart3DShininessMaterialArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DShininessMaterialArchive) ProtoMessage()    {}

func (m *Chart3DShininessMaterialArchive) GetSuper() *Chart3DTexturesMaterialArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *Chart3DShininessMaterialArchive) GetTilings() []*Chart3DImageTextureTilingArchive {
	if m != nil {
		return m.Tilings
	}
	return nil
}

type Chart3DEnvironmentMaterialArchive struct {
	Super            *Chart3DTexturesMaterialArchive         `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	OBSOLETETilings  []*Chart3DBaseImageTextureTilingArchive `protobuf:"bytes,2,rep,name=OBSOLETE_tilings" json:"OBSOLETE_tilings,omitempty"`
	DecalMode        *bool                                   `protobuf:"varint,3,opt,name=decalMode" json:"decalMode,omitempty"`
	Tilings          []*Chart3DImageTextureTilingArchive     `protobuf:"bytes,4,rep,name=tilings" json:"tilings,omitempty"`
	XXX_unrecognized []byte                                  `json:"-"`
}

func (m *Chart3DEnvironmentMaterialArchive) Reset()         { *m = Chart3DEnvironmentMaterialArchive{} }
func (m *Chart3DEnvironmentMaterialArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DEnvironmentMaterialArchive) ProtoMessage()    {}

func (m *Chart3DEnvironmentMaterialArchive) GetSuper() *Chart3DTexturesMaterialArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *Chart3DEnvironmentMaterialArchive) GetOBSOLETETilings() []*Chart3DBaseImageTextureTilingArchive {
	if m != nil {
		return m.OBSOLETETilings
	}
	return nil
}

func (m *Chart3DEnvironmentMaterialArchive) GetDecalMode() bool {
	if m != nil && m.DecalMode != nil {
		return *m.DecalMode
	}
	return false
}

func (m *Chart3DEnvironmentMaterialArchive) GetTilings() []*Chart3DImageTextureTilingArchive {
	if m != nil {
		return m.Tilings
	}
	return nil
}

type Chart3DFixedFunctionLightingModelArchive struct {
	Materials        *Chart3DPhongMaterialPackageArchive `protobuf:"bytes,1,req,name=materials" json:"materials,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *Chart3DFixedFunctionLightingModelArchive) Reset() {
	*m = Chart3DFixedFunctionLightingModelArchive{}
}
func (m *Chart3DFixedFunctionLightingModelArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DFixedFunctionLightingModelArchive) ProtoMessage()    {}

func (m *Chart3DFixedFunctionLightingModelArchive) GetMaterials() *Chart3DPhongMaterialPackageArchive {
	if m != nil {
		return m.Materials
	}
	return nil
}

type Chart3DPhongLightingModelArchive struct {
	Materials        *Chart3DPhongMaterialPackageArchive `protobuf:"bytes,1,req,name=materials" json:"materials,omitempty"`
	XXX_unrecognized []byte                              `json:"-"`
}

func (m *Chart3DPhongLightingModelArchive) Reset()         { *m = Chart3DPhongLightingModelArchive{} }
func (m *Chart3DPhongLightingModelArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DPhongLightingModelArchive) ProtoMessage()    {}

func (m *Chart3DPhongLightingModelArchive) GetMaterials() *Chart3DPhongMaterialPackageArchive {
	if m != nil {
		return m.Materials
	}
	return nil
}

type Chart3DPhongMaterialPackageArchive struct {
	Emissive         *Chart3DEmissiveMaterialArchive  `protobuf:"bytes,1,opt,name=emissive" json:"emissive,omitempty"`
	Diffuse          *Chart3DDiffuseMaterialArchive   `protobuf:"bytes,2,opt,name=diffuse" json:"diffuse,omitempty"`
	Modulate         *Chart3DModulateMaterialArchive  `protobuf:"bytes,3,opt,name=modulate" json:"modulate,omitempty"`
	Specular         *Chart3DSpecularMaterialArchive  `protobuf:"bytes,4,opt,name=specular" json:"specular,omitempty"`
	Shininess        *Chart3DShininessMaterialArchive `protobuf:"bytes,5,opt,name=shininess" json:"shininess,omitempty"`
	XXX_unrecognized []byte                           `json:"-"`
}

func (m *Chart3DPhongMaterialPackageArchive) Reset()         { *m = Chart3DPhongMaterialPackageArchive{} }
func (m *Chart3DPhongMaterialPackageArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DPhongMaterialPackageArchive) ProtoMessage()    {}

func (m *Chart3DPhongMaterialPackageArchive) GetEmissive() *Chart3DEmissiveMaterialArchive {
	if m != nil {
		return m.Emissive
	}
	return nil
}

func (m *Chart3DPhongMaterialPackageArchive) GetDiffuse() *Chart3DDiffuseMaterialArchive {
	if m != nil {
		return m.Diffuse
	}
	return nil
}

func (m *Chart3DPhongMaterialPackageArchive) GetModulate() *Chart3DModulateMaterialArchive {
	if m != nil {
		return m.Modulate
	}
	return nil
}

func (m *Chart3DPhongMaterialPackageArchive) GetSpecular() *Chart3DSpecularMaterialArchive {
	if m != nil {
		return m.Specular
	}
	return nil
}

func (m *Chart3DPhongMaterialPackageArchive) GetShininess() *Chart3DShininessMaterialArchive {
	if m != nil {
		return m.Shininess
	}
	return nil
}

type Chart3DTSPImageDataTextureArchive struct {
	Data               *TSP.DataReference `protobuf:"bytes,3,opt,name=data" json:"data,omitempty"`
	Mipmapdata         *TSP.DataReference `protobuf:"bytes,4,opt,name=mipmapdata" json:"mipmapdata,omitempty"`
	DatabaseData       *TSP.Reference     `protobuf:"bytes,1,opt,name=database_data" json:"database_data,omitempty"`
	DatabaseMipmapdata *TSP.Reference     `protobuf:"bytes,2,opt,name=database_mipmapdata" json:"database_mipmapdata,omitempty"`
	XXX_unrecognized   []byte             `json:"-"`
}

func (m *Chart3DTSPImageDataTextureArchive) Reset()         { *m = Chart3DTSPImageDataTextureArchive{} }
func (m *Chart3DTSPImageDataTextureArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DTSPImageDataTextureArchive) ProtoMessage()    {}

func (m *Chart3DTSPImageDataTextureArchive) GetData() *TSP.DataReference {
	if m != nil {
		return m.Data
	}
	return nil
}

func (m *Chart3DTSPImageDataTextureArchive) GetMipmapdata() *TSP.DataReference {
	if m != nil {
		return m.Mipmapdata
	}
	return nil
}

func (m *Chart3DTSPImageDataTextureArchive) GetDatabaseData() *TSP.Reference {
	if m != nil {
		return m.DatabaseData
	}
	return nil
}

func (m *Chart3DTSPImageDataTextureArchive) GetDatabaseMipmapdata() *TSP.Reference {
	if m != nil {
		return m.DatabaseMipmapdata
	}
	return nil
}

type Chart3DBaseImageTextureTilingArchive struct {
	Scale            *Chart3DVectorArchive `protobuf:"bytes,1,opt,name=scale" json:"scale,omitempty"`
	Rotation         *float32              `protobuf:"fixed32,2,opt,name=rotation" json:"rotation,omitempty"`
	XXX_unrecognized []byte                `json:"-"`
}

func (m *Chart3DBaseImageTextureTilingArchive) Reset()         { *m = Chart3DBaseImageTextureTilingArchive{} }
func (m *Chart3DBaseImageTextureTilingArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DBaseImageTextureTilingArchive) ProtoMessage()    {}

func (m *Chart3DBaseImageTextureTilingArchive) GetScale() *Chart3DVectorArchive {
	if m != nil {
		return m.Scale
	}
	return nil
}

func (m *Chart3DBaseImageTextureTilingArchive) GetRotation() float32 {
	if m != nil && m.Rotation != nil {
		return *m.Rotation
	}
	return 0
}

type Chart3DImageTextureTilingArchive struct {
	Super            *Chart3DBaseImageTextureTilingArchive `protobuf:"bytes,1,req,name=super" json:"super,omitempty"`
	Mode             *TextureTilingMode                    `protobuf:"varint,2,opt,name=mode,enum=TSCH.TextureTilingMode" json:"mode,omitempty"`
	Wrap             *TextureTilingWrap                    `protobuf:"varint,3,opt,name=wrap,enum=TSCH.TextureTilingWrap" json:"wrap,omitempty"`
	Face             *TextureTilingFace                    `protobuf:"varint,4,opt,name=face,enum=TSCH.TextureTilingFace" json:"face,omitempty"`
	Xposition        *TextureTilingXPosition               `protobuf:"varint,5,opt,name=xposition,enum=TSCH.TextureTilingXPosition" json:"xposition,omitempty"`
	Yposition        *TextureTilingYPosition               `protobuf:"varint,6,opt,name=yposition,enum=TSCH.TextureTilingYPosition" json:"yposition,omitempty"`
	Scontinuity      *TextureTilingContinuity              `protobuf:"varint,7,opt,name=scontinuity,enum=TSCH.TextureTilingContinuity" json:"scontinuity,omitempty"`
	Tcontinuity      *TextureTilingContinuity              `protobuf:"varint,8,opt,name=tcontinuity,enum=TSCH.TextureTilingContinuity" json:"tcontinuity,omitempty"`
	Reveal           *bool                                 `protobuf:"varint,9,opt,name=reveal" json:"reveal,omitempty"`
	XXX_unrecognized []byte                                `json:"-"`
}

func (m *Chart3DImageTextureTilingArchive) Reset()         { *m = Chart3DImageTextureTilingArchive{} }
func (m *Chart3DImageTextureTilingArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DImageTextureTilingArchive) ProtoMessage()    {}

func (m *Chart3DImageTextureTilingArchive) GetSuper() *Chart3DBaseImageTextureTilingArchive {
	if m != nil {
		return m.Super
	}
	return nil
}

func (m *Chart3DImageTextureTilingArchive) GetMode() TextureTilingMode {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return TextureTilingMode_textureTilingModeNone
}

func (m *Chart3DImageTextureTilingArchive) GetWrap() TextureTilingWrap {
	if m != nil && m.Wrap != nil {
		return *m.Wrap
	}
	return TextureTilingWrap_textureTilingWrapProjected
}

func (m *Chart3DImageTextureTilingArchive) GetFace() TextureTilingFace {
	if m != nil && m.Face != nil {
		return *m.Face
	}
	return TextureTilingFace_textureTilingFaceAll
}

func (m *Chart3DImageTextureTilingArchive) GetXposition() TextureTilingXPosition {
	if m != nil && m.Xposition != nil {
		return *m.Xposition
	}
	return TextureTilingXPosition_textureTilingXPositionLeft
}

func (m *Chart3DImageTextureTilingArchive) GetYposition() TextureTilingYPosition {
	if m != nil && m.Yposition != nil {
		return *m.Yposition
	}
	return TextureTilingYPosition_textureTilingYPositionTop
}

func (m *Chart3DImageTextureTilingArchive) GetScontinuity() TextureTilingContinuity {
	if m != nil && m.Scontinuity != nil {
		return *m.Scontinuity
	}
	return TextureTilingContinuity_textureTilingContinuityNone
}

func (m *Chart3DImageTextureTilingArchive) GetTcontinuity() TextureTilingContinuity {
	if m != nil && m.Tcontinuity != nil {
		return *m.Tcontinuity
	}
	return TextureTilingContinuity_textureTilingContinuityNone
}

func (m *Chart3DImageTextureTilingArchive) GetReveal() bool {
	if m != nil && m.Reveal != nil {
		return *m.Reveal
	}
	return false
}

type Chart3DVectorArchive struct {
	X                *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	Z                *float32 `protobuf:"fixed32,3,req,name=z" json:"z,omitempty"`
	W                *float32 `protobuf:"fixed32,4,req,name=w" json:"w,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Chart3DVectorArchive) Reset()         { *m = Chart3DVectorArchive{} }
func (m *Chart3DVectorArchive) String() string { return proto.CompactTextString(m) }
func (*Chart3DVectorArchive) ProtoMessage()    {}

func (m *Chart3DVectorArchive) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Chart3DVectorArchive) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

func (m *Chart3DVectorArchive) GetZ() float32 {
	if m != nil && m.Z != nil {
		return *m.Z
	}
	return 0
}

func (m *Chart3DVectorArchive) GetW() float32 {
	if m != nil && m.W != nil {
		return *m.W
	}
	return 0
}

var E_Fill3D = &proto.ExtensionDesc{
	ExtendedType:  (*TSD.FillArchive)(nil),
	ExtensionType: (*Chart3DFillArchive)(nil),
	Field:         100,
	Name:          "TSCH.fill3d",
	Tag:           "bytes,100,opt,name=fill3d",
}

func init() {
	proto.RegisterEnum("TSCH.TextureTilingMode", TextureTilingMode_name, TextureTilingMode_value)
	proto.RegisterEnum("TSCH.TextureTilingFace", TextureTilingFace_name, TextureTilingFace_value)
	proto.RegisterEnum("TSCH.TextureTilingWrap", TextureTilingWrap_name, TextureTilingWrap_value)
	proto.RegisterEnum("TSCH.TextureTilingXPosition", TextureTilingXPosition_name, TextureTilingXPosition_value)
	proto.RegisterEnum("TSCH.TextureTilingYPosition", TextureTilingYPosition_name, TextureTilingYPosition_value)
	proto.RegisterEnum("TSCH.TextureTilingContinuity", TextureTilingContinuity_name, TextureTilingContinuity_value)
	proto.RegisterEnum("TSCH.FillPropertyType", FillPropertyType_name, FillPropertyType_value)
	proto.RegisterExtension(E_Fill3D)
}
