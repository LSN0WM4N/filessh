package theme

type IconType = rune

func CharacterToDIsplay(c *Config, name string) rune {
	realName := name
	if !c.UFT8() {
		realName += "_RETRO"
	}

	switch realName {
	case "FOLDER":
		return FOLDER_ICON
	case "FILE":
		return FILE_ICON
	case "IMAGE":
		return IMAGE_ICON
	case "AUDIO":
		return AUDIO_ICON
	case "VIDEO":
		return VIDEO_ICON
	case "PDF":
		fallthrough
	case "DOCUMENT":
		return PDF_ICON
	case "BACK":
		return BACK_ICON

	case "FOLDER_RETRO":
		return FOLDER_ICON_RETRO
	case "FILE_RETRO":
		fallthrough
	case "IMAGE_RETRO":
		fallthrough
	case "AUDIO_RETRO":
		fallthrough
	case "VIDEO_RETRO":
		fallthrough
	case "PDF_RETRO":
		fallthrough
	case "DOCUMENT_RETRO":
		return FILE_ICON_RETRO
	case "BACK_RETRO":
		return BACK_ICON_RETRO

	default:
		return ' '
	}
}

const (
	FOLDER_ICON IconType = '\U0001F5C0'
	FILE_ICON   IconType = '\U0001F5B9'
	BACK_ICON   IconType = '\U0001F8A8'
	IMAGE_ICON  IconType = '\U0001F5BC'
	AUDIO_ICON  IconType = '\U0000266B'
	VIDEO_ICON  IconType = '\U000025B6'
	PDF_ICON    IconType = '\U000025A4'

	// Backward compatibility
	FOLDER_ICON_RETRO IconType = '+'
	FILE_ICON_RETRO   IconType = '$'
	BACK_ICON_RETRO   IconType = '<'
)
