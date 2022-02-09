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

// UpdateNodeRequest update node request
//
// swagger:model UpdateNodeRequest
type UpdateNodeRequest struct {

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

// Validate validates this update node request
func (m *UpdateNodeRequest) Validate(formats strfmt.Registry) error {
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

func (m *UpdateNodeRequest) validateAnswers(formats strfmt.Registry) error {

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

func (m *UpdateNodeRequest) validateAssetType(formats strfmt.Registry) error {

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

func (m *UpdateNodeRequest) validateCreatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.CreatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("createdAt", "body", "date-time", m.CreatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validateCriticality(formats strfmt.Registry) error {

	if err := validate.Required("criticality", "body", m.Criticality); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validateIndustry(formats strfmt.Registry) error {

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

var updateNodeRequestTypeInspectionTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["numeric","single_choice","multi_choice"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNodeRequestTypeInspectionTypePropEnum = append(updateNodeRequestTypeInspectionTypePropEnum, v)
	}
}

const (

	// UpdateNodeRequestInspectionTypeNumeric captures enum value "numeric"
	UpdateNodeRequestInspectionTypeNumeric string = "numeric"

	// UpdateNodeRequestInspectionTypeSingleChoice captures enum value "single_choice"
	UpdateNodeRequestInspectionTypeSingleChoice string = "single_choice"

	// UpdateNodeRequestInspectionTypeMultiChoice captures enum value "multi_choice"
	UpdateNodeRequestInspectionTypeMultiChoice string = "multi_choice"
)

// prop value enum
func (m *UpdateNodeRequest) validateInspectionTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNodeRequestTypeInspectionTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateNodeRequest) validateInspectionType(formats strfmt.Registry) error {

	if err := validate.Required("inspectionType", "body", m.InspectionType); err != nil {
		return err
	}

	// value enum
	if err := m.validateInspectionTypeEnum("inspectionType", "body", *m.InspectionType); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validateLabel(formats strfmt.Registry) error {

	if err := validate.Required("label", "body", m.Label); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validateLocked(formats strfmt.Registry) error {

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

func (m *UpdateNodeRequest) validateLubricant(formats strfmt.Registry) error {

	if err := validate.Required("lubricant", "body", m.Lubricant); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validateLubricantQuantity(formats strfmt.Registry) error {

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

var updateNodeRequestTypeLubricationActivityAssetStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["must_be_on","must_be_off"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNodeRequestTypeLubricationActivityAssetStatePropEnum = append(updateNodeRequestTypeLubricationActivityAssetStatePropEnum, v)
	}
}

const (

	// UpdateNodeRequestLubricationActivityAssetStateMustBeOn captures enum value "must_be_on"
	UpdateNodeRequestLubricationActivityAssetStateMustBeOn string = "must_be_on"

	// UpdateNodeRequestLubricationActivityAssetStateMustBeOff captures enum value "must_be_off"
	UpdateNodeRequestLubricationActivityAssetStateMustBeOff string = "must_be_off"
)

// prop value enum
func (m *UpdateNodeRequest) validateLubricationActivityAssetStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNodeRequestTypeLubricationActivityAssetStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateNodeRequest) validateLubricationActivityAssetState(formats strfmt.Registry) error {

	if swag.IsZero(m.LubricationActivityAssetState) { // not required
		return nil
	}

	// value enum
	if err := m.validateLubricationActivityAssetStateEnum("lubricationActivityAssetState", "body", m.LubricationActivityAssetState); err != nil {
		return err
	}

	return nil
}

var updateNodeRequestTypeMeasurementTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["displacement","acceleration","velocity","temperature","dc_gap","ampl_phase","box","speed","envelope_2","envelope_3","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNodeRequestTypeMeasurementTypePropEnum = append(updateNodeRequestTypeMeasurementTypePropEnum, v)
	}
}

const (

	// UpdateNodeRequestMeasurementTypeDisplacement captures enum value "displacement"
	UpdateNodeRequestMeasurementTypeDisplacement string = "displacement"

	// UpdateNodeRequestMeasurementTypeAcceleration captures enum value "acceleration"
	UpdateNodeRequestMeasurementTypeAcceleration string = "acceleration"

	// UpdateNodeRequestMeasurementTypeVelocity captures enum value "velocity"
	UpdateNodeRequestMeasurementTypeVelocity string = "velocity"

	// UpdateNodeRequestMeasurementTypeTemperature captures enum value "temperature"
	UpdateNodeRequestMeasurementTypeTemperature string = "temperature"

	// UpdateNodeRequestMeasurementTypeDcGap captures enum value "dc_gap"
	UpdateNodeRequestMeasurementTypeDcGap string = "dc_gap"

	// UpdateNodeRequestMeasurementTypeAmplPhase captures enum value "ampl_phase"
	UpdateNodeRequestMeasurementTypeAmplPhase string = "ampl_phase"

	// UpdateNodeRequestMeasurementTypeBox captures enum value "box"
	UpdateNodeRequestMeasurementTypeBox string = "box"

	// UpdateNodeRequestMeasurementTypeSpeed captures enum value "speed"
	UpdateNodeRequestMeasurementTypeSpeed string = "speed"

	// UpdateNodeRequestMeasurementTypeEnvelope2 captures enum value "envelope_2"
	UpdateNodeRequestMeasurementTypeEnvelope2 string = "envelope_2"

	// UpdateNodeRequestMeasurementTypeEnvelope3 captures enum value "envelope_3"
	UpdateNodeRequestMeasurementTypeEnvelope3 string = "envelope_3"

	// UpdateNodeRequestMeasurementTypeUnknown captures enum value "unknown"
	UpdateNodeRequestMeasurementTypeUnknown string = "unknown"
)

// prop value enum
func (m *UpdateNodeRequest) validateMeasurementTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNodeRequestTypeMeasurementTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateNodeRequest) validateMeasurementType(formats strfmt.Registry) error {

	if err := validate.Required("measurementType", "body", m.MeasurementType); err != nil {
		return err
	}

	// value enum
	if err := m.validateMeasurementTypeEnum("measurementType", "body", *m.MeasurementType); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validateMetadata(formats strfmt.Registry) error {

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

var updateNodeRequestTypeOrientationPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["axial","radial","radial90","horizontal","vertical","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNodeRequestTypeOrientationPropEnum = append(updateNodeRequestTypeOrientationPropEnum, v)
	}
}

const (

	// UpdateNodeRequestOrientationAxial captures enum value "axial"
	UpdateNodeRequestOrientationAxial string = "axial"

	// UpdateNodeRequestOrientationRadial captures enum value "radial"
	UpdateNodeRequestOrientationRadial string = "radial"

	// UpdateNodeRequestOrientationRadial90 captures enum value "radial90"
	UpdateNodeRequestOrientationRadial90 string = "radial90"

	// UpdateNodeRequestOrientationHorizontal captures enum value "horizontal"
	UpdateNodeRequestOrientationHorizontal string = "horizontal"

	// UpdateNodeRequestOrientationVertical captures enum value "vertical"
	UpdateNodeRequestOrientationVertical string = "vertical"

	// UpdateNodeRequestOrientationUnknown captures enum value "unknown"
	UpdateNodeRequestOrientationUnknown string = "unknown"
)

// prop value enum
func (m *UpdateNodeRequest) validateOrientationEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNodeRequestTypeOrientationPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateNodeRequest) validateOrientation(formats strfmt.Registry) error {

	if err := validate.Required("orientation", "body", m.Orientation); err != nil {
		return err
	}

	// value enum
	if err := m.validateOrientationEnum("orientation", "body", *m.Orientation); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validateOrigin(formats strfmt.Registry) error {

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

func (m *UpdateNodeRequest) validateParentID(formats strfmt.Registry) error {

	if err := validate.Required("parentId", "body", m.ParentID); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validatePosition(formats strfmt.Registry) error {

	if swag.IsZero(m.Position) { // not required
		return nil
	}

	if err := validate.MinimumInt("position", "body", int64(*m.Position), 0, false); err != nil {
		return err
	}

	return nil
}

var updateNodeRequestTypeSubtypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["root","company","site","plant","ship","system","functional_location","asset","measurement_location","measurement_point","inspection_point","lubrication_point"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNodeRequestTypeSubtypePropEnum = append(updateNodeRequestTypeSubtypePropEnum, v)
	}
}

const (

	// UpdateNodeRequestSubtypeRoot captures enum value "root"
	UpdateNodeRequestSubtypeRoot string = "root"

	// UpdateNodeRequestSubtypeCompany captures enum value "company"
	UpdateNodeRequestSubtypeCompany string = "company"

	// UpdateNodeRequestSubtypeSite captures enum value "site"
	UpdateNodeRequestSubtypeSite string = "site"

	// UpdateNodeRequestSubtypePlant captures enum value "plant"
	UpdateNodeRequestSubtypePlant string = "plant"

	// UpdateNodeRequestSubtypeShip captures enum value "ship"
	UpdateNodeRequestSubtypeShip string = "ship"

	// UpdateNodeRequestSubtypeSystem captures enum value "system"
	UpdateNodeRequestSubtypeSystem string = "system"

	// UpdateNodeRequestSubtypeFunctionalLocation captures enum value "functional_location"
	UpdateNodeRequestSubtypeFunctionalLocation string = "functional_location"

	// UpdateNodeRequestSubtypeAsset captures enum value "asset"
	UpdateNodeRequestSubtypeAsset string = "asset"

	// UpdateNodeRequestSubtypeMeasurementLocation captures enum value "measurement_location"
	UpdateNodeRequestSubtypeMeasurementLocation string = "measurement_location"

	// UpdateNodeRequestSubtypeMeasurementPoint captures enum value "measurement_point"
	UpdateNodeRequestSubtypeMeasurementPoint string = "measurement_point"

	// UpdateNodeRequestSubtypeInspectionPoint captures enum value "inspection_point"
	UpdateNodeRequestSubtypeInspectionPoint string = "inspection_point"

	// UpdateNodeRequestSubtypeLubricationPoint captures enum value "lubrication_point"
	UpdateNodeRequestSubtypeLubricationPoint string = "lubrication_point"
)

// prop value enum
func (m *UpdateNodeRequest) validateSubtypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNodeRequestTypeSubtypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateNodeRequest) validateSubtype(formats strfmt.Registry) error {

	if err := validate.Required("subtype", "body", m.Subtype); err != nil {
		return err
	}

	// value enum
	if err := m.validateSubtypeEnum("subtype", "body", *m.Subtype); err != nil {
		return err
	}

	return nil
}

var updateNodeRequestTypeTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["root","company","site","plant","system","functional_location","asset","measurement_location","measurement_point","inspection_point","lubrication_point","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNodeRequestTypeTypePropEnum = append(updateNodeRequestTypeTypePropEnum, v)
	}
}

const (

	// UpdateNodeRequestTypeRoot captures enum value "root"
	UpdateNodeRequestTypeRoot string = "root"

	// UpdateNodeRequestTypeCompany captures enum value "company"
	UpdateNodeRequestTypeCompany string = "company"

	// UpdateNodeRequestTypeSite captures enum value "site"
	UpdateNodeRequestTypeSite string = "site"

	// UpdateNodeRequestTypePlant captures enum value "plant"
	UpdateNodeRequestTypePlant string = "plant"

	// UpdateNodeRequestTypeSystem captures enum value "system"
	UpdateNodeRequestTypeSystem string = "system"

	// UpdateNodeRequestTypeFunctionalLocation captures enum value "functional_location"
	UpdateNodeRequestTypeFunctionalLocation string = "functional_location"

	// UpdateNodeRequestTypeAsset captures enum value "asset"
	UpdateNodeRequestTypeAsset string = "asset"

	// UpdateNodeRequestTypeMeasurementLocation captures enum value "measurement_location"
	UpdateNodeRequestTypeMeasurementLocation string = "measurement_location"

	// UpdateNodeRequestTypeMeasurementPoint captures enum value "measurement_point"
	UpdateNodeRequestTypeMeasurementPoint string = "measurement_point"

	// UpdateNodeRequestTypeInspectionPoint captures enum value "inspection_point"
	UpdateNodeRequestTypeInspectionPoint string = "inspection_point"

	// UpdateNodeRequestTypeLubricationPoint captures enum value "lubrication_point"
	UpdateNodeRequestTypeLubricationPoint string = "lubrication_point"

	// UpdateNodeRequestTypeUnknown captures enum value "unknown"
	UpdateNodeRequestTypeUnknown string = "unknown"
)

// prop value enum
func (m *UpdateNodeRequest) validateTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNodeRequestTypeTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateNodeRequest) validateType(formats strfmt.Registry) error {

	if err := validate.Required("type", "body", m.Type); err != nil {
		return err
	}

	// value enum
	if err := m.validateTypeEnum("type", "body", *m.Type); err != nil {
		return err
	}

	return nil
}

func (m *UpdateNodeRequest) validateUpdatedAt(formats strfmt.Registry) error {

	if swag.IsZero(m.UpdatedAt) { // not required
		return nil
	}

	if err := validate.FormatOf("updatedAt", "body", "date-time", m.UpdatedAt.String(), formats); err != nil {
		return err
	}

	return nil
}

var updateNodeRequestTypeVisualizationTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["visualization_none","visualization_circular_gauge","visualization_level_gauge"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		updateNodeRequestTypeVisualizationTypePropEnum = append(updateNodeRequestTypeVisualizationTypePropEnum, v)
	}
}

const (

	// UpdateNodeRequestVisualizationTypeVisualizationNone captures enum value "visualization_none"
	UpdateNodeRequestVisualizationTypeVisualizationNone string = "visualization_none"

	// UpdateNodeRequestVisualizationTypeVisualizationCircularGauge captures enum value "visualization_circular_gauge"
	UpdateNodeRequestVisualizationTypeVisualizationCircularGauge string = "visualization_circular_gauge"

	// UpdateNodeRequestVisualizationTypeVisualizationLevelGauge captures enum value "visualization_level_gauge"
	UpdateNodeRequestVisualizationTypeVisualizationLevelGauge string = "visualization_level_gauge"
)

// prop value enum
func (m *UpdateNodeRequest) validateVisualizationTypeEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, updateNodeRequestTypeVisualizationTypePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *UpdateNodeRequest) validateVisualizationType(formats strfmt.Registry) error {

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
func (m *UpdateNodeRequest) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *UpdateNodeRequest) UnmarshalBinary(b []byte) error {
	var res UpdateNodeRequest
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}