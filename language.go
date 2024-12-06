package main

var langMap = map[string]map[string]map[string]string{
	"changePasswordPage": {
		"en": {
			"pageTitle":        "Change Password",
			"oldPasswordLabel": "Old Password:",
			"newPasswordLabel": "New Password:",
			"oldPasswordHint":  "Enter your Old Password here.\t",
			"newPasswordHint":  "Enter your New Password here.\t",
		},
		"fr": {
			"pageTitle":        "changer le mot de passe",
			"oldPasswordLabel": "Ancien mot de passe:",
			"newPasswordLabel": "Nouveau mot de passe:",
			"oldPasswordHint":  "Entrez votre ancien mot de passe ici.\t",
			"newPasswordHint":  "Entrez votre nouveau mot de passe ici.\t",
		},
		"ar": {
			"pageTitle":        "تغير كلمة السر",
			"oldPasswordLabel": "كلمة السر القديمة :",
			"newPasswordLabel": "كلمة السر الجديدة :",
			"oldPasswordHint":  "ادخل كلمة السر الحالية هنا        ",
			"newPasswordHint":  "ادخل كلمة السر الجديدة هنا        ",
		},
	},
	"homePage": {
		"en": {
			"pageTitle":              "Home Page",
			"searchInputPlaceHolder": "search",
			"noFileSelected":         "Select a file from the list",
		},
		"fr": {},
		"ar": {},
	},
	"loginPage": {
		"en": {
			"wrongPassword": "wrong password",
		},
		"fr": {
			"wrongPassword": "FR_wrong password",
		},
		"ar": {
			"wrongPassword": "AR_wrong password",
		},
	},
	"settingsPage": {
		"en": {
			"pageTitle":                  "Settings",
			"themeDark":                  "Dark Theme",
			"themeLight":                 "Light Theme",
			"themeLabel":                 "Theme:",
			"languageLabel":              "Language:",
			"changePasswordButton":       "Change Password",
			"changeSaveFolderButton":     "change save folder",
			"saveFolderPathLabelDeafult": "Deafult",
			"englishLanguage":            "English",
			"frenchLanguage":             "Francais",
			"arabicLanguage":             "العربية",
		},
		"fr": {
			"pageTitle":                  "FR_Settings",
			"themeDark":                  "FR_Dark Theme",
			"themeLight":                 "FR_Light Theme",
			"themeLabel":                 "FR_Theme:",
			"languageLabel":              "FR_Language:",
			"changePasswordButton":       "FR_Change Password",
			"changeSaveFolderButton":     "FR_change save folder",
			"saveFolderPathLabelDeafult": "FR_Deafult",
			"englishLanguage":            "English",
			"frenchLanguage":             "Francais",
			"arabicLanguage":             "العربية",
		},
		"ar": {
			"pageTitle":                  "الإعدادات",
			"themeDark":                  "مظلم",
			"themeLight":                 "مضيء",
			"themeLabel":                 "الوضع:",
			"languageLabel":              "اللغة:",
			"changePasswordButton":       "تغير كلمة السر",
			"changeSaveFolderButton":     "غير مجلد حفظ الملفات",
			"saveFolderPathLabelDeafult": "AR_Deafult",
			"englishLanguage":            "English",
			"frenchLanguage":             "Francais",
			"arabicLanguage":             "العربية",
		},
	},
	"addFilePage": {
		"en": {
			"pageTitle":           "Add a new file",
			"fileNameLabel":       "File Name:",
			"fileNameHintText":    "Enter a File Name here, this will be used when searching.\t",
			"openFileLabel":       "File:",
			"openFileButton":      "Open a file",
			"openFileHintText":    "Choose the file that you want to add.\t",
			"neededDateLabel":     "Needed Date:",
			"neededDateButton":    "Select a date",
			"neededDateHintText":  "Choose when You gonna need the File (optional).\t",
			"startDateLabel":      "Start Date:",
			"startDateButton":     "Select a date",
			"startDateHintText":   "This makres the date when the file was created (optional).\t",
			"endDateLabel":        "End Date:",
			"endDateButton":       "Select a date",
			"endDateHintText":     "This makres the date when the file will serve no purpse (optional).\t",
			"locationLabel":       "Location:",
			"locationHintText":    "Enter where the file is stored in real life (optional).\t",
			"descriptionLabel":    "Description:",
			"descriptionHintText": "Enter a description for the file (optional).\t",
			"formSubmit":          "Submit",
		},
		"fr": {
			"pageTitle": "",
		},
		"ar": {
			"pageTitle": "",
		},
	},
	"calendarDialog": {
		"en": {
			"title":  "Select a date",
			"cancel": "Cancel",
		},
		"fr": {
			"pageTitle": "",
		},
		"ar": {
			"pageTitle": "",
		},
	},
}
