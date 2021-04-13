/* Host definition */
package nagios

import "time"

type Info struct {
	Created         time.Time //created=1589796613
	Version         string    //version=4.4.5
	LastUpdateCheck time.Time //last_update_check=1589716104
	UpdateAvailable int       //update_available=0
	LastVersion     string    //last_version=4.4.5
	NewVersion      string    //new_version=4.4.5
}

func NewInfoFromMap(m map[string]string) (Info, error) {
	info := Info{}
	// created
	createdTimestamp, err := parseIntInMap(m, "created")
	if err != nil {
		return info, err
	}
	info.Created = time.Unix(createdTimestamp, 0)
	// last update check
	lastUpdateCheckTimestamp, err := parseIntInMap(m, "last_update_check")
	if err != nil {
		return info, err
	}
	info.LastUpdateCheck = time.Unix(lastUpdateCheckTimestamp, 0)
	// UpdateAvailable
	updateAvailable, err := parseIntInMap(m, "update_available")
	if err != nil {
		return info, err
	}
	info.UpdateAvailable = int(updateAvailable)
	// version
	info.Version = m["version"]
	info.LastVersion = m["last_version"]
	info.NewVersion = m["new_version"]
	return info, nil
}
