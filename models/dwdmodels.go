// Package dwdmodels provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/oapi-codegen/oapi-codegen/v2 version v2.3.0 DO NOT EDIT.
package dwdmodels

import (
	"encoding/json"

	openapi_types "github.com/oapi-codegen/runtime/types"
)

// CROWDMeldung defines model for CROWDMeldung.
type CROWDMeldung struct {
	End               *int64 `json:"end,omitempty"`
	HighestSeverities *[]struct {
		Auspraegung *string `json:"auspraegung,omitempty"`
		Category    *string `json:"category,omitempty"`
	} `json:"highestSeverities,omitempty"`
	Meldungen *[]struct {
		Auspraegung     *string                   `json:"auspraegung,omitempty"`
		Category        *string                   `json:"category,omitempty"`
		Lat             *string                   `json:"lat,omitempty"`
		Lon             *string                   `json:"lon,omitempty"`
		MeldungId       *int32                    `json:"meldungId,omitempty"`
		Place           *string                   `json:"place,omitempty"`
		Timestamp       *int64                    `json:"timestamp,omitempty"`
		ZusatzAttribute *[]map[string]interface{} `json:"zusatzAttribute,omitempty"`
	} `json:"meldungen,omitempty"`
	Start *int64 `json:"start,omitempty"`
}

// Error defines model for Error.
type Error struct {
	Error     *string `json:"error,omitempty"`
	Message   *string `json:"message,omitempty"`
	Path      *string `json:"path,omitempty"`
	Status    *int32  `json:"status,omitempty"`
	Timestamp *int64  `json:"timestamp,omitempty"`
}

// GemeindeWarnings defines model for GemeindeWarnings.
type GemeindeWarnings struct {
	BinnenSee *struct {
		N209901000 *[]struct {
			Bn              *bool   `json:"bn,omitempty"`
			Description     *string `json:"description,omitempty"`
			DescriptionText *string `json:"descriptionText,omitempty"`
			Event           *string `json:"event,omitempty"`
			Headline        *string `json:"headline,omitempty"`
			Instruction     *string `json:"instruction,omitempty"`
			Level           *int32  `json:"level,omitempty"`
			Start           *int64  `json:"start,omitempty"`
			Type            *int32  `json:"type,omitempty"`
		} `json:"209901000,omitempty"`
	} `json:"binnenSee,omitempty"`
	Time     *int64 `json:"time,omitempty"`
	Warnings *[]struct {
		Bn              *bool   `json:"bn,omitempty"`
		Description     *string `json:"description,omitempty"`
		DescriptionText *string `json:"descriptionText,omitempty"`
		End             *int64  `json:"end,omitempty"`
		Event           *string `json:"event,omitempty"`
		HeadLine        *string `json:"headLine,omitempty"`
		IsVorabinfo     *bool   `json:"isVorabinfo,omitempty"`
		Level           *int32  `json:"level,omitempty"`
		Regions         *[]struct {
			Polygon   *[]float32 `json:"polygon,omitempty"`
			Triangles *[]int32   `json:"triangles,omitempty"`
		} `json:"regions,omitempty"`
		Start *int64                    `json:"start,omitempty"`
		Type  *int32                    `json:"type,omitempty"`
		Urls  *[]map[string]interface{} `json:"urls,omitempty"`
	} `json:"warnings,omitempty"`
}

// StationOverview defines model for StationOverview.
type StationOverview struct {
	N10865 *struct {
		Days *[]struct {
			DayDate        *openapi_types.Date `json:"dayDate,omitempty"`
			Icon           *int32              `json:"icon,omitempty"`
			Icon1          *string             `json:"icon1"`
			Icon2          *string             `json:"icon2"`
			Precipitation  *int32              `json:"precipitation,omitempty"`
			StationId      *string             `json:"stationId"`
			Sunshine       *int32              `json:"sunshine,omitempty"`
			TemperatureMax *int32              `json:"temperatureMax,omitempty"`
			TemperatureMin *int32              `json:"temperatureMin,omitempty"`
			WindDirection  *int32              `json:"windDirection,omitempty"`
			WindGust       *int32              `json:"windGust,omitempty"`
			WindSpeed      *int32              `json:"windSpeed,omitempty"`
		} `json:"days,omitempty"`
		Forecast1 *struct {
			Icon                         *[]int32 `json:"icon,omitempty"`
			PrecipitationProbablity      *string  `json:"precipitationProbablity"`
			PrecipitationProbablityIndex *string  `json:"precipitationProbablityIndex"`
			PrecipitationTotal           *[]int32 `json:"precipitationTotal,omitempty"`
			Start                        *int64   `json:"start,omitempty"`
			StationId                    *string  `json:"stationId,omitempty"`
			Temperature                  *[]int32 `json:"temperature,omitempty"`
			TemperatureStd               *[]int32 `json:"temperatureStd,omitempty"`
			TimeStep                     *int32   `json:"timeStep,omitempty"`
			WindDirection                *string  `json:"windDirection"`
			WindGust                     *string  `json:"windGust"`
			WindSpeed                    *string  `json:"windSpeed"`
		} `json:"forecast1,omitempty"`
		Forecast2 *struct {
			Icon                         *[]int32   `json:"icon,omitempty"`
			PrecipitationProbablity      *string    `json:"precipitationProbablity"`
			PrecipitationProbablityIndex *string    `json:"precipitationProbablityIndex"`
			PrecipitationTotal           *[]int32   `json:"precipitationTotal,omitempty"`
			Start                        *int64     `json:"start,omitempty"`
			StationId                    *string    `json:"stationId,omitempty"`
			Temperature                  *[]float32 `json:"temperature,omitempty"`
			TemperatureStd               *[]float32 `json:"temperatureStd,omitempty"`
			TimeStep                     *int32     `json:"timeStep,omitempty"`
			WindDirection                *string    `json:"windDirection"`
			WindGust                     *string    `json:"windGust"`
			WindSpeed                    *string    `json:"windSpeed"`
		} `json:"forecast2,omitempty"`
		ForecastStart      *string                   `json:"forecastStart"`
		ThreeHourSummaries *string                   `json:"threeHourSummaries"`
		Warnings           *[]map[string]interface{} `json:"warnings,omitempty"`
	} `json:"10865,omitempty"`
}

// WarningCoast defines model for WarningCoast.
type WarningCoast struct {
	Time     *int64 `json:"time,omitempty"`
	Warnings *struct {
		N501000007 *[]struct {
			Bn              *bool   `json:"bn,omitempty"`
			Description     *string `json:"description,omitempty"`
			DescriptionText *string `json:"descriptionText,omitempty"`
			Event           *string `json:"event,omitempty"`
			Headline        *string `json:"headline,omitempty"`
			Instruction     *string `json:"instruction,omitempty"`
			Level           *int32  `json:"level,omitempty"`
			Type            *int32  `json:"type,omitempty"`
		} `json:"501000007,omitempty"`
	} `json:"warnings,omitempty"`
}

// WarningNowcast defines model for WarningNowcast.
type WarningNowcast struct {
	BinnenSee *string `json:"binnenSee"`
	Time      *int64  `json:"time,omitempty"`
	Warnings  *[]struct {
		Bn              *bool   `json:"bn,omitempty"`
		Description     *string `json:"description,omitempty"`
		DescriptionText *string `json:"descriptionText,omitempty"`
		End             *int64  `json:"end,omitempty"`
		Event           *string `json:"event,omitempty"`
		HeadLine        *string `json:"headLine,omitempty"`
		Instruction     *string `json:"instruction"`
		IsVorabinfo     *bool   `json:"isVorabinfo,omitempty"`
		Level           *int32  `json:"level,omitempty"`
		Regions         *[]struct {
			Polygon   *[]float32 `json:"polygon,omitempty"`
			Triangles *[]int32   `json:"triangles,omitempty"`
		} `json:"regions,omitempty"`
		Start  *int64     `json:"start,omitempty"`
		States *[]float32 `json:"states,omitempty"`
		Type   *int32     `json:"type,omitempty"`
		Urls   *[]string  `json:"urls,omitempty"`
	} `json:"warnings,omitempty"`
}

// GetStationOverviewExtendedParams defines parameters for GetStationOverviewExtended.
type GetStationOverviewExtendedParams struct {
	// StationIds Beim Parameter stationsIds handelt es sich um die Stationskennungen des DWD.
	// Die Liste der Stationskennungen kann z.B. [hier](https://www.dwd.de/DE/leistungen/klimadatendeutschland/stationsliste.html)
	// eingesehen werden.
	StationIds *[]GetStationOverviewExtendedParams_StationIds_Item `form:"stationIds,omitempty" json:"stationIds,omitempty"`
}

// GetStationOverviewExtendedParamsStationIds0 defines parameters for GetStationOverviewExtended.
type GetStationOverviewExtendedParamsStationIds0 = string

// GetStationOverviewExtendedParamsStationIds1 defines parameters for GetStationOverviewExtended.
type GetStationOverviewExtendedParamsStationIds1 = int

// GetStationOverviewExtendedParams_StationIds_Item defines parameters for GetStationOverviewExtended.
type GetStationOverviewExtendedParams_StationIds_Item struct {
	union json.RawMessage
}