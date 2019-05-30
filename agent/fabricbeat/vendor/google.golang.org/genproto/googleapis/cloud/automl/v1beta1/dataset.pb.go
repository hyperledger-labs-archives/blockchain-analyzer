// Code generated by protoc-gen-go. DO NOT EDIT.
// source: google/cloud/automl/v1beta1/dataset.proto

package automl

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
	_ "google.golang.org/genproto/googleapis/api/annotations"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// A workspace for solving a single, particular machine learning (ML) problem.
// A workspace contains examples that may be annotated.
type Dataset struct {
	// Required.
	// The dataset metadata that is specific to the problem type.
	//
	// Types that are valid to be assigned to DatasetMetadata:
	//	*Dataset_TranslationDatasetMetadata
	//	*Dataset_ImageClassificationDatasetMetadata
	//	*Dataset_TextClassificationDatasetMetadata
	//	*Dataset_ImageObjectDetectionDatasetMetadata
	//	*Dataset_VideoClassificationDatasetMetadata
	//	*Dataset_VideoObjectTrackingDatasetMetadata
	//	*Dataset_TextExtractionDatasetMetadata
	//	*Dataset_TextSentimentDatasetMetadata
	//	*Dataset_TablesDatasetMetadata
	DatasetMetadata isDataset_DatasetMetadata `protobuf_oneof:"dataset_metadata"`
	// Output only. The resource name of the dataset.
	// Form: `projects/{project_id}/locations/{location_id}/datasets/{dataset_id}`
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Required. The name of the dataset to show in the interface. The name can be
	// up to 32 characters long and can consist only of ASCII Latin letters A-Z
	// and a-z, underscores
	// (_), and ASCII digits 0-9.
	DisplayName string `protobuf:"bytes,2,opt,name=display_name,json=displayName,proto3" json:"display_name,omitempty"`
	// User-provided description of the dataset. The description can be up to
	// 25000 characters long.
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	// Output only. The number of examples in the dataset.
	ExampleCount int32 `protobuf:"varint,21,opt,name=example_count,json=exampleCount,proto3" json:"example_count,omitempty"`
	// Output only. Timestamp when this dataset was created.
	CreateTime *timestamp.Timestamp `protobuf:"bytes,14,opt,name=create_time,json=createTime,proto3" json:"create_time,omitempty"`
	// Used to perform consistent read-modify-write updates. If not set, a blind
	// "overwrite" update happens.
	Etag                 string   `protobuf:"bytes,17,opt,name=etag,proto3" json:"etag,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dataset) Reset()         { *m = Dataset{} }
func (m *Dataset) String() string { return proto.CompactTextString(m) }
func (*Dataset) ProtoMessage()    {}
func (*Dataset) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f2b1dc66a1e92da, []int{0}
}

func (m *Dataset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dataset.Unmarshal(m, b)
}
func (m *Dataset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dataset.Marshal(b, m, deterministic)
}
func (m *Dataset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dataset.Merge(m, src)
}
func (m *Dataset) XXX_Size() int {
	return xxx_messageInfo_Dataset.Size(m)
}
func (m *Dataset) XXX_DiscardUnknown() {
	xxx_messageInfo_Dataset.DiscardUnknown(m)
}

var xxx_messageInfo_Dataset proto.InternalMessageInfo

type isDataset_DatasetMetadata interface {
	isDataset_DatasetMetadata()
}

type Dataset_TranslationDatasetMetadata struct {
	TranslationDatasetMetadata *TranslationDatasetMetadata `protobuf:"bytes,23,opt,name=translation_dataset_metadata,json=translationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageClassificationDatasetMetadata struct {
	ImageClassificationDatasetMetadata *ImageClassificationDatasetMetadata `protobuf:"bytes,24,opt,name=image_classification_dataset_metadata,json=imageClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_TextClassificationDatasetMetadata struct {
	TextClassificationDatasetMetadata *TextClassificationDatasetMetadata `protobuf:"bytes,25,opt,name=text_classification_dataset_metadata,json=textClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_ImageObjectDetectionDatasetMetadata struct {
	ImageObjectDetectionDatasetMetadata *ImageObjectDetectionDatasetMetadata `protobuf:"bytes,26,opt,name=image_object_detection_dataset_metadata,json=imageObjectDetectionDatasetMetadata,proto3,oneof"`
}

type Dataset_VideoClassificationDatasetMetadata struct {
	VideoClassificationDatasetMetadata *VideoClassificationDatasetMetadata `protobuf:"bytes,31,opt,name=video_classification_dataset_metadata,json=videoClassificationDatasetMetadata,proto3,oneof"`
}

type Dataset_VideoObjectTrackingDatasetMetadata struct {
	VideoObjectTrackingDatasetMetadata *VideoObjectTrackingDatasetMetadata `protobuf:"bytes,29,opt,name=video_object_tracking_dataset_metadata,json=videoObjectTrackingDatasetMetadata,proto3,oneof"`
}

type Dataset_TextExtractionDatasetMetadata struct {
	TextExtractionDatasetMetadata *TextExtractionDatasetMetadata `protobuf:"bytes,28,opt,name=text_extraction_dataset_metadata,json=textExtractionDatasetMetadata,proto3,oneof"`
}

type Dataset_TextSentimentDatasetMetadata struct {
	TextSentimentDatasetMetadata *TextSentimentDatasetMetadata `protobuf:"bytes,30,opt,name=text_sentiment_dataset_metadata,json=textSentimentDatasetMetadata,proto3,oneof"`
}

type Dataset_TablesDatasetMetadata struct {
	TablesDatasetMetadata *TablesDatasetMetadata `protobuf:"bytes,33,opt,name=tables_dataset_metadata,json=tablesDatasetMetadata,proto3,oneof"`
}

func (*Dataset_TranslationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_ImageObjectDetectionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_VideoClassificationDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_VideoObjectTrackingDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextExtractionDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TextSentimentDatasetMetadata) isDataset_DatasetMetadata() {}

func (*Dataset_TablesDatasetMetadata) isDataset_DatasetMetadata() {}

func (m *Dataset) GetDatasetMetadata() isDataset_DatasetMetadata {
	if m != nil {
		return m.DatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTranslationDatasetMetadata() *TranslationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TranslationDatasetMetadata); ok {
		return x.TranslationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageClassificationDatasetMetadata() *ImageClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageClassificationDatasetMetadata); ok {
		return x.ImageClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextClassificationDatasetMetadata() *TextClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextClassificationDatasetMetadata); ok {
		return x.TextClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetImageObjectDetectionDatasetMetadata() *ImageObjectDetectionDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_ImageObjectDetectionDatasetMetadata); ok {
		return x.ImageObjectDetectionDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetVideoClassificationDatasetMetadata() *VideoClassificationDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_VideoClassificationDatasetMetadata); ok {
		return x.VideoClassificationDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetVideoObjectTrackingDatasetMetadata() *VideoObjectTrackingDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_VideoObjectTrackingDatasetMetadata); ok {
		return x.VideoObjectTrackingDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextExtractionDatasetMetadata() *TextExtractionDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextExtractionDatasetMetadata); ok {
		return x.TextExtractionDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTextSentimentDatasetMetadata() *TextSentimentDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TextSentimentDatasetMetadata); ok {
		return x.TextSentimentDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetTablesDatasetMetadata() *TablesDatasetMetadata {
	if x, ok := m.GetDatasetMetadata().(*Dataset_TablesDatasetMetadata); ok {
		return x.TablesDatasetMetadata
	}
	return nil
}

func (m *Dataset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dataset) GetDisplayName() string {
	if m != nil {
		return m.DisplayName
	}
	return ""
}

func (m *Dataset) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Dataset) GetExampleCount() int32 {
	if m != nil {
		return m.ExampleCount
	}
	return 0
}

func (m *Dataset) GetCreateTime() *timestamp.Timestamp {
	if m != nil {
		return m.CreateTime
	}
	return nil
}

func (m *Dataset) GetEtag() string {
	if m != nil {
		return m.Etag
	}
	return ""
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Dataset) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Dataset_TranslationDatasetMetadata)(nil),
		(*Dataset_ImageClassificationDatasetMetadata)(nil),
		(*Dataset_TextClassificationDatasetMetadata)(nil),
		(*Dataset_ImageObjectDetectionDatasetMetadata)(nil),
		(*Dataset_VideoClassificationDatasetMetadata)(nil),
		(*Dataset_VideoObjectTrackingDatasetMetadata)(nil),
		(*Dataset_TextExtractionDatasetMetadata)(nil),
		(*Dataset_TextSentimentDatasetMetadata)(nil),
		(*Dataset_TablesDatasetMetadata)(nil),
	}
}

func init() {
	proto.RegisterType((*Dataset)(nil), "google.cloud.automl.v1beta1.Dataset")
}

func init() {
	proto.RegisterFile("google/cloud/automl/v1beta1/dataset.proto", fileDescriptor_1f2b1dc66a1e92da)
}

var fileDescriptor_1f2b1dc66a1e92da = []byte{
	// 633 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x95, 0xdd, 0x6e, 0xd3, 0x3c,
	0x18, 0xc7, 0x5f, 0xef, 0xe5, 0x43, 0xb8, 0x03, 0x81, 0xa5, 0x69, 0x21, 0xeb, 0x68, 0xbb, 0xc1,
	0x56, 0x0e, 0x48, 0xb4, 0x71, 0x80, 0x60, 0x12, 0xb0, 0x75, 0x08, 0x38, 0x18, 0xa0, 0x52, 0xed,
	0x00, 0x55, 0x8a, 0xdc, 0xc4, 0x8b, 0x0c, 0x49, 0x1c, 0x25, 0x4f, 0xab, 0x22, 0xce, 0x10, 0x5c,
	0x00, 0x12, 0x48, 0x5c, 0x13, 0x77, 0xc0, 0xdd, 0x20, 0x7f, 0x94, 0x01, 0x69, 0x9d, 0x9d, 0xa5,
	0x7e, 0x7e, 0xfd, 0xfb, 0xd7, 0xbf, 0x23, 0x17, 0xdf, 0x8e, 0x85, 0x88, 0x13, 0xe6, 0x87, 0x89,
	0x18, 0x47, 0x3e, 0x1d, 0x83, 0x48, 0x13, 0x7f, 0xb2, 0x33, 0x62, 0x40, 0x77, 0xfc, 0x88, 0x02,
	0x2d, 0x19, 0x78, 0x79, 0x21, 0x40, 0x90, 0x35, 0x8d, 0x7a, 0x0a, 0xf5, 0x34, 0xea, 0x19, 0xd4,
	0xdd, 0xb6, 0xe5, 0xf0, 0x94, 0xc6, 0x4c, 0xa7, 0xb8, 0x5d, 0x1b, 0x08, 0x74, 0x94, 0xb0, 0xd2,
	0x90, 0x5b, 0x56, 0x92, 0x4d, 0x8d, 0x97, 0x7b, 0xc7, 0xca, 0x15, 0x34, 0x2b, 0x13, 0x0a, 0x5c,
	0x64, 0x06, 0xb7, 0x9a, 0x4e, 0x78, 0xc4, 0x84, 0x01, 0x5b, 0x06, 0x54, 0x9f, 0x46, 0xe3, 0x13,
	0x1f, 0x78, 0xca, 0x4a, 0xa0, 0x69, 0x6e, 0x80, 0xa6, 0x01, 0x68, 0xce, 0x7d, 0x9a, 0x65, 0x02,
	0xd4, 0x36, 0x46, 0x7f, 0xe3, 0x67, 0x03, 0x5f, 0x3c, 0xd4, 0x05, 0x92, 0x0f, 0xb8, 0xf9, 0x87,
	0x48, 0x60, 0x7a, 0x0d, 0x52, 0x06, 0x54, 0x3e, 0x3b, 0xab, 0x6d, 0xd4, 0x6d, 0xec, 0xde, 0xf3,
	0x2c, 0x0d, 0x7b, 0x83, 0xd3, 0x00, 0x13, 0x7b, 0x64, 0xbe, 0xfe, 0xec, 0xbf, 0xbe, 0x0b, 0x0b,
	0xa7, 0xe4, 0x2b, 0xc2, 0xb7, 0xd4, 0x09, 0x04, 0x61, 0x42, 0xcb, 0x92, 0x9f, 0xf0, 0x70, 0x81,
	0x86, 0xa3, 0x34, 0x1e, 0x59, 0x35, 0x9e, 0xcb, 0xa4, 0xde, 0x5f, 0x41, 0x55, 0x9d, 0x0d, 0x5e,
	0x4b, 0x91, 0x2f, 0x08, 0xdf, 0x94, 0xa7, 0x58, 0x6b, 0x75, 0x5d, 0x59, 0x3d, 0xb4, 0x97, 0xc3,
	0xa6, 0x50, 0x27, 0xd5, 0x81, 0x3a, 0x88, 0x7c, 0x47, 0x78, 0x5b, 0x57, 0x25, 0x46, 0x6f, 0x59,
	0x08, 0x41, 0xc4, 0x80, 0x85, 0xf3, 0xb5, 0x5c, 0xa5, 0xf5, 0xb8, 0xbe, 0xac, 0x97, 0x2a, 0xea,
	0x70, 0x96, 0x54, 0x15, 0xdb, 0xe4, 0xf5, 0x98, 0x3a, 0x45, 0xf5, 0x76, 0xd6, 0xf6, 0xd5, 0x3a,
	0xc3, 0x29, 0x1e, 0xcb, 0xa4, 0xda, 0x53, 0x9c, 0xd4, 0x52, 0xe4, 0x1b, 0xc2, 0x5b, 0x5a, 0xcb,
	0x34, 0x06, 0x05, 0x0d, 0xdf, 0xf1, 0x2c, 0xae, 0x7a, 0xad, 0x9f, 0xd5, 0x4b, 0x37, 0x31, 0x30,
	0x41, 0x8b, 0xbc, 0xac, 0x14, 0xf9, 0x8c, 0x70, 0x5b, 0xbd, 0x5d, 0x6c, 0x2a, 0x8d, 0xe6, 0x37,
	0xd5, 0x54, 0x46, 0x0f, 0x6a, 0xdf, 0xac, 0x27, 0xbf, 0x33, 0xaa, 0x32, 0xeb, 0x60, 0x03, 0xc8,
	0x47, 0x84, 0x5b, 0xca, 0xa3, 0x64, 0x99, 0xbc, 0x3f, 0x32, 0xa8, 0x6a, 0xdc, 0x50, 0x1a, 0xf7,
	0x6b, 0x35, 0x5e, 0xcf, 0x22, 0xaa, 0x16, 0x4d, 0xb0, 0xcc, 0x49, 0x82, 0x57, 0xf5, 0xcd, 0x5a,
	0xdd, 0xbb, 0xa3, 0xf6, 0xde, 0xb5, 0xef, 0xad, 0xbe, 0x5b, 0xdd, 0x74, 0x05, 0xe6, 0x0d, 0x08,
	0xc1, 0xe7, 0x32, 0x9a, 0x32, 0x07, 0xb5, 0x51, 0xf7, 0x52, 0x5f, 0x3d, 0x93, 0x0e, 0x5e, 0x8e,
	0x78, 0x99, 0x27, 0xf4, 0x7d, 0xa0, 0x66, 0x4b, 0x6a, 0xd6, 0x30, 0x6b, 0x2f, 0x24, 0xd2, 0xc6,
	0x8d, 0x88, 0x95, 0x61, 0xc1, 0x73, 0xd9, 0xa3, 0xf3, 0xbf, 0x21, 0x4e, 0x97, 0xc8, 0x26, 0xbe,
	0xcc, 0xa6, 0x34, 0xcd, 0x13, 0x16, 0x84, 0x62, 0x9c, 0x81, 0xb3, 0xd2, 0x46, 0xdd, 0xf3, 0xfd,
	0x65, 0xb3, 0xd8, 0x93, 0x6b, 0x64, 0x0f, 0x37, 0xc2, 0x82, 0x51, 0x60, 0x81, 0xec, 0xc2, 0xb9,
	0xa2, 0x7e, 0x9f, 0x3b, 0xfb, 0x7d, 0xb3, 0xbb, 0xdc, 0x1b, 0xcc, 0xee, 0xf2, 0x3e, 0xd6, 0xb8,
	0x5c, 0x90, 0xea, 0x0c, 0x68, 0xec, 0x5c, 0xd3, 0xea, 0xf2, 0xf9, 0x80, 0xe0, 0xab, 0xff, 0xb6,
	0x76, 0xf0, 0x09, 0xe1, 0x56, 0x28, 0x52, 0x5b, 0x6b, 0xaf, 0xd0, 0x9b, 0x7d, 0x33, 0x8e, 0x45,
	0x42, 0xb3, 0xd8, 0x13, 0x45, 0xec, 0xc7, 0x2c, 0x53, 0x0a, 0xbe, 0x1e, 0xd1, 0x9c, 0x97, 0x73,
	0xff, 0x88, 0xf6, 0xf4, 0xc7, 0x1f, 0x4b, 0x6b, 0x4f, 0x15, 0x38, 0xec, 0x49, 0x68, 0xb8, 0x3f,
	0x06, 0x71, 0x94, 0x0c, 0x8f, 0x35, 0x34, 0xba, 0xa0, 0xb2, 0xee, 0xfe, 0x0a, 0x00, 0x00, 0xff,
	0xff, 0x9f, 0x88, 0xc9, 0xd2, 0xc5, 0x07, 0x00, 0x00,
}
