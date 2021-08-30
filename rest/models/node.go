// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"strconv"

	"github.com/SKF/go-utility/v2/uuid"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// Node node
//
// swagger:model Node
type Node struct {

	// Possible answers to a single choice or multi choice inspection.
	Answers []*Answer `json:"answers"`

	// Only valid for asset type nodes
	AssetType *AssetType `json:"assetType,omitempty"`

	// Which country the node is in. Only valid for site type node
	Country string `json:"country,omitempty"`

	// Timestamp of node creation
	// Format: date-time
	CreatedAt strfmt.DateTime `json:"createdAt,omitempty"`

	// Only valid for asset type nodes
	// Required: true
	Criticality *string `json:"criticality"`

	// Type of device used to take measurements on this point. Only valid for measurement point type nodes
	DadType string `json:"dadType,omitempty"`

	// Description of the node
	Description string `json:"description,omitempty"`

	// ID of node, as a UUID
	ID uuid.UUID `json:"id,omitempty"`

	// Industry segment of this node. Only valid for site type node
	Industry *Industry `json:"industry,omitempty"`

	// Type of value to record.
	// Required: true
	// Enum: [numeric single_choice multi_choice]
	InspectionType *string `json:"inspectionType"`

	// Descriptive name of the node
	// Required: true
	Label *string `json:"label"`

	// Locked indicates if this node is locked for editing by someone
	Locked *Lock `json:"locked,omitempty"`

	// Type of lubricant used. Only valid for lubrication point type nodes
	// Required: true
	Lubricant *string `json:"lubricant"`

	// Amount of lubricant to be used. Only valid for lubrication point type nodes
	LubricantQuantity *NaturalQuantity `json:"lubricantQuantity,omitempty"`

	// Instruction for lubrication activity. Only valid for lubrication point type nodes
	LubricateInstructions string `json:"lubricateInstructions,omitempty"`

	// ActivityAssetState the asset should be in during the lubrication activity. Only valid for lubrication point type nodes
	// Enum: [must_be_on must_be_off]
	LubricationActivityAssetState string `json:"lubricationActivityAssetState,omitempty"`

	// Only valid for asset type nodes
	Manufacturer string `json:"manufacturer,omitempty"`

	// Type of measurement. Only valid for measurement point type nodes
	// Required: true
	// Enum: [displacement acceleration velocity temperature dc_gap ampl_phase box speed envelope_2 envelope_3 unknown]
	MeasurementType *string `json:"measurementType"`

	// Metadata with keys and optional values
	Metadata NodeMetaData `json:"metadata,omitempty"`

	// Only valid for asset type nodes
	Model string `json:"model,omitempty"`

	// Orientation of measurement. Only valid for measurement point type nodes
	// Required: true
	// Enum: [axial radial radial90 horizontal vertical unknown]
	Orientation *string `json:"orientation"`

	// Origin of node, if imported from another system
	Origin *Origin `json:"origin,omitempty"`

	// ID of parent node, as a UUID
	// Required: true
	ParentID *string `json:"parentId"`

	// Relative position of node in the Enlight Centre UI
	// Minimum: 0
	Position *int64 `json:"position,omitempty"`

	// Which postal code the site has. Only valid for site type node
	PostalCode string `json:"postalCode,omitempty"`

	// Only valid for asset type nodes
	SerialNumber string `json:"serialNumber,omitempty"`

	// Subtype of node
	// Required: true
	// Enum: [root company site plant ship system functional_location asset measurement_location measurement_point inspection_point lubrication_point]
	Subtype *string `json:"subtype"`

	// Type of node
	// Required: true
	// Enum: [root company site plant system functional_location asset measurement_location measurement_point inspection_point lubrication_point unknown]
	Type *string `json:"type"`

	// Unit of value recorded. Only valid for inspection points of type `numeric`.
	Unit string `json:"unit,omitempty"`

	// Timestamp of last node update
	// Format: date-time
	UpdatedAt strfmt.DateTime `json:"updatedAt,omitempty"`

	// Maximum value used when visualizing the value of the point on a scale. Only valid for inspection points of type `numeric`.
	VisualizationMaxValue float64 `json:"visualizationMaxValue,omitempty"`

	// Minimum value used when visualizing the value of the point on a scale. Only valid for inspection points of type `numeric`.
	VisualizationMinValue float64 `json:"visualizationMinValue,omitempty"`

	// Type of visualization in Enlight Centre. Only valid for inspection points of type `numeric`.
	// Enum: [visualization_none visualization_circular_gauge visualization_level_gauge]
	VisualizationType string `json:"visualizationType,omitempty"`
}

// Validate validates this node
func (m *Node) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAnswers(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateAssetType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCriticality(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIndustry(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateInspectionType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLabel(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLocked(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLubricant(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLubricantQuantity(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLubricationActivityAssetState(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMeasurementType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMetadata(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrientation(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateOrigin(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateParentID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePosition(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateSubtype(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpdatedAt(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVisualizationType(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *Node) validateAnswers(formats strfmt.Registry) error {

	if swag.IsZero(m.Answers) { // not required
		return nil
	}

	for i := 0; i < len(m.Answers); i++ {
		if swag.IsZero(m.Answers[i]) { // not required
			continue
		}

		if m.Answers[i] != nil {
			if err := m.Answers[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("answers" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *Node) validateAssetType(formats strfmt.Registry) error {

	if swag.IsZero(m.AssetType) { // not required
		return nil
	}

	if m.AssetType != nil {
		if err := m.AssetType.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("assetType")
			}
			return err
		}
	}

	return nil
}

func (m *Node) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateCriticality(formats strfmt.Registry) error {

	if err := validate.Required("criticality", "body", m.Criticality); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateIndustry(formats strfmt.Registry) error {

	if swag.IsZero(m.Industry) { // not required
		return nil
	}

	if m.Industry != nil {
		if err := m.Industry.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("industry")
			}
			return err
		}
	}

	return nil
}

var nodeTypeInspectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["numeric","single_choice","multi_choice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeInspectionTypePropEnum = append(nodeTypeInspectionTypePropEnum, v)
	}
}

const (

	// NodeInspectionTypeNumeric captures enum value "numeric"
	NodeInspectionTypeNumeric string = "numeric"

	// NodeInspectionTypeSingleChoice captures enum value "single_choice"
	NodeInspectionTypeSingleChoice string = "single_choice"

	// NodeInspectionTypeMultiChoice captures enum value "multi_choice"
	NodeInspectionTypeMultiChoice string = "multi_choice"
)

// prop value enum
func (m *Node) validateInspectionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeInspectionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateInspectionType(formats strfmt.Registry) error {

	if err := validate.Required("inspectionType", "body", m.InspectionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateInspectionTypeEnum("inspectionType", "body", *m.InspectionType); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateLocked(formats strfmt.Registry) error {

	if swag.IsZero(m.Locked) { // not required
		return nil
	}

	if m.Locked != nil {
		if err := m.Locked.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("locked")
			}
			return err
		}
	}

	return nil
}

func (m *Node) validateLubricant(formats strfmt.Registry) error {

	if err := validate.Required("lubricant", "body", m.Lubricant); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateLubricantQuantity(formats strfmt.Registry) error {

	if swag.IsZero(m.LubricantQuantity) { // not required
		return nil
	}

	if m.LubricantQuantity != nil {
		if err := m.LubricantQuantity.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("lubricantQuantity")
			}
			return err
		}
	}

	return nil
}

var nodeTypeLubricationActivityAssetStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["must_be_on","must_be_off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeLubricationActivityAssetStatePropEnum = append(nodeTypeLubricationActivityAssetStatePropEnum, v)
	}
}

const (

	// NodeLubricationActivityAssetStateMustBeOn captures enum value "must_be_on"
	NodeLubricationActivityAssetStateMustBeOn string = "must_be_on"

	// NodeLubricationActivityAssetStateMustBeOff captures enum value "must_be_off"
	NodeLubricationActivityAssetStateMustBeOff string = "must_be_off"
)

// prop value enum
func (m *Node) validateLubricationActivityAssetStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeLubricationActivityAssetStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateLubricationActivityAssetState(formats strfmt.Registry) error {

	if swag.IsZero(m.LubricationActivityAssetState) { // not required
		return nil
	}

	// value enum
	if err := m.validateLubricationActivityAssetStateEnum("lubricationActivityAssetState", "body", m.LubricationActivityAssetState); err != nil {
		return err
	}

	return nil
}

var nodeTypeMeasurementTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["displacement","acceleration","velocity","temperature","dc_gap","ampl_phase","box","speed","envelope_2","envelope_3","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeMeasurementTypePropEnum = append(nodeTypeMeasurementTypePropEnum, v)
	}
}

const (

	// NodeMeasurementTypeDisplacement captures enum value "displacement"
	NodeMeasurementTypeDisplacement string = "displacement"

	// NodeMeasurementTypeAcceleration captures enum value "acceleration"
	NodeMeasurementTypeAcceleration string = "acceleration"

	// NodeMeasurementTypeVelocity captures enum value "velocity"
	NodeMeasurementTypeVelocity string = "velocity"

	// NodeMeasurementTypeTemperature captures enum value "temperature"
	NodeMeasurementTypeTemperature string = "temperature"

	// NodeMeasurementTypeDcGap captures enum value "dc_gap"
	NodeMeasurementTypeDcGap string = "dc_gap"

	// NodeMeasurementTypeAmplPhase captures enum value "ampl_phase"
	NodeMeasurementTypeAmplPhase string = "ampl_phase"

	// NodeMeasurementTypeBox captures enum value "box"
	NodeMeasurementTypeBox string = "box"

	// NodeMeasurementTypeSpeed captures enum value "speed"
	NodeMeasurementTypeSpeed string = "speed"

	// NodeMeasurementTypeEnvelope2 captures enum value "envelope_2"
	NodeMeasurementTypeEnvelope2 string = "envelope_2"

	// NodeMeasurementTypeEnvelope3 captures enum value "envelope_3"
	NodeMeasurementTypeEnvelope3 string = "envelope_3"

	// NodeMeasurementTypeUnknown captures enum value "unknown"
	NodeMeasurementTypeUnknown string = "unknown"
)

// prop value enum
func (m *Node) validateMeasurementTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeMeasurementTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateMeasurementType(formats strfmt.Registry) error {

	if err := validate.Required("measurementType", "body", m.MeasurementType); err != nil {
		return err
	}

	// value enum
	if err := m.validateMeasurementTypeEnum("measurementType", "body", *m.MeasurementType); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateMetadata(formats strfmt.Registry) error {

	if swag.IsZero(m.Metadata) { // not required
		return nil
	}

	if err := m.Metadata.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("metadata")
		}
		return err
	}

	return nil
}

var nodeTypeOrientationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["axial","radial","radial90","horizontal","vertical","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeOrientationPropEnum = append(nodeTypeOrientationPropEnum, v)
	}
}

const (

	// NodeOrientationAxial captures enum value "axial"
	NodeOrientationAxial string = "axial"

	// NodeOrientationRadial captures enum value "radial"
	NodeOrientationRadial string = "radial"

	// NodeOrientationRadial90 captures enum value "radial90"
	NodeOrientationRadial90 string = "radial90"

	// NodeOrientationHorizontal captures enum value "horizontal"
	NodeOrientationHorizontal string = "horizontal"

	// NodeOrientationVertical captures enum value "vertical"
	NodeOrientationVertical string = "vertical"

	// NodeOrientationUnknown captures enum value "unknown"
	NodeOrientationUnknown string = "unknown"
)

// prop value enum
func (m *Node) validateOrientationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeOrientationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateOrientation(formats strfmt.Registry) error {

	if err := validate.Required("orientation", "body", m.Orientation); err != nil {
		return err
	}

	// value enum
	if err := m.validateOrientationEnum("orientation", "body", *m.Orientation); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateOrigin(formats strfmt.Registry) error {

	if swag.IsZero(m.Origin) { // not required
		return nil
	}

	if m.Origin != nil {
		if err := m.Origin.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("origin")
			}
			return err
		}
	}

	return nil
}

func (m *Node) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parentId", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

func (m *Node) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", int64(*m.Position), 0, false); err != nil {
		return err
	}

	return nil
}

var nodeTypeSubtypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["root","company","site","plant","ship","system","functional_location","asset","measurement_location","measurement_point","inspection_point","lubrication_point"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeSubtypePropEnum = append(nodeTypeSubtypePropEnum, v)
	}
}

const (

	// NodeSubtypeRoot captures enum value "root"
	NodeSubtypeRoot string = "root"

	// NodeSubtypeCompany captures enum value "company"
	NodeSubtypeCompany string = "company"

	// NodeSubtypeSite captures enum value "site"
	NodeSubtypeSite string = "site"

	// NodeSubtypePlant captures enum value "plant"
	NodeSubtypePlant string = "plant"

	// NodeSubtypeShip captures enum value "ship"
	NodeSubtypeShip string = "ship"

	// NodeSubtypeSystem captures enum value "system"
	NodeSubtypeSystem string = "system"

	// NodeSubtypeFunctionalLocation captures enum value "functional_location"
	NodeSubtypeFunctionalLocation string = "functional_location"

	// NodeSubtypeAsset captures enum value "asset"
	NodeSubtypeAsset string = "asset"

	// NodeSubtypeMeasurementLocation captures enum value "measurement_location"
	NodeSubtypeMeasurementLocation string = "measurement_location"

	// NodeSubtypeMeasurementPoint captures enum value "measurement_point"
	NodeSubtypeMeasurementPoint string = "measurement_point"

	// NodeSubtypeInspectionPoint captures enum value "inspection_point"
	NodeSubtypeInspectionPoint string = "inspection_point"

	// NodeSubtypeLubricationPoint captures enum value "lubrication_point"
	NodeSubtypeLubricationPoint string = "lubrication_point"
)

// prop value enum
func (m *Node) validateSubtypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeSubtypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateSubtype(formats strfmt.Registry) error {

	if err := validate.Required("subtype", "body", m.Subtype); err != nil {
		return err
	}

	// value enum
	if err := m.validateSubtypeEnum("subtype", "body", *m.Subtype); err != nil {
		return err
	}

	return nil
}

var nodeTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["root","company","site","plant","system","functional_location","asset","measurement_location","measurement_point","inspection_point","lubrication_point","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeTypePropEnum = append(nodeTypeTypePropEnum, v)
	}
}

const (

	// NodeTypeRoot captures enum value "root"
	NodeTypeRoot string = "root"

	// NodeTypeCompany captures enum value "company"
	NodeTypeCompany string = "company"

	// NodeTypeSite captures enum value "site"
	NodeTypeSite string = "site"

	// NodeTypePlant captures enum value "plant"
	NodeTypePlant string = "plant"

	// NodeTypeSystem captures enum value "system"
	NodeTypeSystem string = "system"

	// NodeTypeFunctionalLocation captures enum value "functional_location"
	NodeTypeFunctionalLocation string = "functional_location"

	// NodeTypeAsset captures enum value "asset"
	NodeTypeAsset string = "asset"

	// NodeTypeMeasurementLocation captures enum value "measurement_location"
	NodeTypeMeasurementLocation string = "measurement_location"

	// NodeTypeMeasurementPoint captures enum value "measurement_point"
	NodeTypeMeasurementPoint string = "measurement_point"

	// NodeTypeInspectionPoint captures enum value "inspection_point"
	NodeTypeInspectionPoint string = "inspection_point"

	// NodeTypeLubricationPoint captures enum value "lubrication_point"
	NodeTypeLubricationPoint string = "lubrication_point"

	// NodeTypeUnknown captures enum value "unknown"
	NodeTypeUnknown string = "unknown"
)

// prop value enum
func (m *Node) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *Node) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var nodeTypeVisualizationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["visualization_none","visualization_circular_gauge","visualization_level_gauge"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nodeTypeVisualizationTypePropEnum = append(nodeTypeVisualizationTypePropEnum, v)
	}
}

const (

	// NodeVisualizationTypeVisualizationNone captures enum value "visualization_none"
	NodeVisualizationTypeVisualizationNone string = "visualization_none"

	// NodeVisualizationTypeVisualizationCircularGauge captures enum value "visualization_circular_gauge"
	NodeVisualizationTypeVisualizationCircularGauge string = "visualization_circular_gauge"

	// NodeVisualizationTypeVisualizationLevelGauge captures enum value "visualization_level_gauge"
	NodeVisualizationTypeVisualizationLevelGauge string = "visualization_level_gauge"
)

// prop value enum
func (m *Node) validateVisualizationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, nodeTypeVisualizationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *Node) validateVisualizationType(formats strfmt.Registry) error {

	if swag.IsZero(m.VisualizationType) { // not required
		return nil
	}

	// value enum
	if err := m.validateVisualizationTypeEnum("visualizationType", "body", m.VisualizationType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *Node) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Node) UnmarshalBinary(b []byte) error {
	var res Node
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
