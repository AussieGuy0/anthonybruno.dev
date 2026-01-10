package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

const (
	dateFormat = "2006-01-02"
)

func WriteActivities(activities []Activity, writeFolder string) error {
	for _, activity := range activities {
		err := writeActivity(&activity, writeFolder)
		if err != nil {
			return err
		}
	}
	return nil
}

func writeActivity(activity *Activity, path string) error {
	str := generateActivityString(activity)
	filename, err := generateFilename(activity)
	if err != nil {
		return err
	}

	filePath := path + "/" + filename
	log.Println("Writing to " + filePath)
	err = os.WriteFile(filePath, []byte(str), 0644)
	if err != nil {
		return err
	}
	return nil
}

func generateFilename(activity *Activity) (string, error) {
	startDate, err := activity.StartDateTime()
	if err != nil {
		return "", err
	}
	// Use activity type and ID for filename to ensure uniqueness
	activityType := strings.ToLower(activity.Type)

	return startDate.Format(dateFormat) + "-" + activityType + "-" + strconv.FormatInt(activity.Id, 10) + ".md", nil
}

func generateActivityString(activity *Activity) string {
	var sb strings.Builder

	sb.WriteString("---\n")
	writeKey(&sb, "title", activity.Type)
	writeKey(&sb, "activity_type", activity.Type)
	writeKey(&sb, "distance_km", fmt.Sprintf("%.2f", activity.DistanceKm()))
	writeKey(&sb, "moving_time", strconv.Itoa(activity.MovingTime))
	writeKey(&sb, "moving_time_formatted", activity.MovingTimeFormatted())
	writeKey(&sb, "elapsed_time", strconv.Itoa(activity.ElapsedTime))
	writeKey(&sb, "elevation_gain", fmt.Sprintf("%.1f", activity.TotalElevationGain))
	writeKey(&sb, "average_speed_kmh", fmt.Sprintf("%.2f", activity.AverageSpeedKmh()))
	writeKey(&sb, "pace_min_per_km", "\""+activity.PaceMinPerKm()+"\"")
	writeKey(&sb, "layout", "activity")
	sb.WriteString("---\n")

	return sb.String()
}

func writeKey(sb *strings.Builder, key string, value string) {
	sb.WriteString(key)
	sb.WriteString(": ")
	sb.WriteString(value)
	sb.WriteString("\n")
}
