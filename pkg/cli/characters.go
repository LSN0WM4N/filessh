package cli

type IconType = rune

const (
	FOLDER_ICON IconType = '\U0001F5C0'
	FILE_ICON   IconType = '\U0001F5B9'
	BACK_ICON   IconType = '\U0001F8A8'

	// Backward compatibility
	FOLDER_ICON_RETRO IconType = '+'
	FILE_ICON_RETRO   IconType = '$'
	BACK_ICON_RETRO   IconType = '<'
)
