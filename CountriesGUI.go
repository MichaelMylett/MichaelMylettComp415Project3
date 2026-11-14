package main

import (
	"log"
	"strconv"
	"strings"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

var ( //this file also uses rest example as a template a lot
	countries               []Country
	commonNameDisplay       *widget.Entry
	officialNameDisplay     *widget.Entry
	languagesDisplay        *widget.Entry
	flagDisplay             *canvas.Image //was originally widget, changed to fit flag png in
	landlockedStatusDisplay *widget.Entry
	populationDisplay       *widget.Entry
	currencyTypesDisplay    *widget.Entry
	countryRegionDisplay    *widget.Entry
)

func buildGUI(countryData []Country) fyne.Window {
	guiApp := app.New()
	countries = countryData
	//var filteredCountries []Country
	mainWindow := guiApp.NewWindow("Project 3 GUI")
	directions := widget.NewLabel("Select a region")
	directions.Alignment = fyne.TextAlignCenter
	directions.TextStyle = fyne.TextStyle{Bold: true}
	var countryList *widget.List

	regionsList := []string{"Africa", "Americas", "Asia", "Europe", "Oceania"}
	regionSelect := widget.NewSelect(regionsList, func(selected string) { //had to look up some of this, but was not copied from AI
		regionFilteredCountries, err := getCountryData(selected)
		if err != nil {
			log.Fatal(err)
		} else {
			countries = regionFilteredCountries //was filteredCountries
			countryList.Refresh()               //had to look up refresh and unselect all
			countryList.UnselectAll()
		}
	})
	regionSelectorPanel := container.New(layout.NewVBoxLayout(), directions, regionSelect)

	countryList = widget.NewList(func() int { return len(countries) }, //was copied from rest demo
		func() fyne.CanvasObject {
			return widget.NewLabel("Placeholder")
		},
		func(i widget.ListItemID, selected fyne.CanvasObject) {
			selected.(*widget.Label).SetText(countries[i].Names.Official)
		})
	countryList.OnSelected = countrySelected
	detailsPanel := makeDetailsPanel()
	detailsPanelWithScroll := container.NewVScroll(detailsPanel) //had to look up how to do this, not AI'd
	mainLayout := container.NewBorder(regionSelectorPanel, nil, nil, detailsPanelWithScroll, countryList)
	mainWindow.SetContent(mainLayout)
	mainWindow.Resize(fyne.NewSize(1900, 1080)) //assuming you meant 1920 x 1080 for window size but did 1900 just in case
	return mainWindow
}

func countrySelected(i widget.ListItemID) {
	//official name
	countryOfficialName := countries[i].Names.Official
	if len(countryOfficialName) == 0 {
		officialNameDisplay.SetText("Official Name Not Found")
		return
	} else {
		officialNameDisplay.SetText(countryOfficialName)
	}
	//common name
	countryCommonName := countries[i].Names.Common
	if len(countryCommonName) == 0 {
		commonNameDisplay.SetText("Common Name Not Found")
	} else {
		commonNameDisplay.SetText(countryCommonName)
	}
	//languages
	languagesdata := countries[i].Languages
	if languagesdata == nil {
		languagesDisplay.SetText("Languages not available")
		return
	} else {
		var languages = strings.Builder{}
		for _, language := range languagesdata {
			languages.WriteString(language.Name + "\n")
		}
		languagesDisplay.SetText(languages.String())
	}
	//flag
	flagURL := countries[i].Flag.URLPNG
	if len(flagURL) == 0 {
		flagDisplay = nil
	} else {
		image, err := fyne.LoadResourceFromURLString(flagURL)
		if err != nil {
			log.Fatal(err)
		} else {
			flagDisplay.Resource = image
			flagDisplay.Refresh()
		}
	}
	//landlocked?
	if countries[i].Landlocked {
		landlockedStatusDisplay.SetText("This country is landlocked.")
	} else {
		landlockedStatusDisplay.SetText("This country has ocean access.")
	}
	//population
	countryPopulation := countries[i].Population
	if countryPopulation <= 0 {
		populationDisplay.SetText("The population is either 0 or unavailable at this time.")
	} else {
		populationStr := strconv.FormatInt(int64(countryPopulation), 10) //looked up FormatInt command
		populationDisplay.SetText(populationStr)
	}
	//currency types
	currencyData := countries[i].Currencies
	if currencyData == nil {
		currencyTypesDisplay.SetText("Currencies not available")
		return
	} else {
		var currencies = strings.Builder{}
		for _, currency := range currencyData {
			currencies.WriteString(currency.Name + "\n")
		}
		currencyTypesDisplay.SetText(currencies.String())
	}
	//region
	countryRegionDisplay.SetText(countries[i].Region)
}

func makeDetailsPanel() *fyne.Container {
	detailsPanel := container.New(layout.NewVBoxLayout())
	countryRegionPanel := container.New(layout.NewVBoxLayout())
	officialNamePanel := container.New(layout.NewVBoxLayout())
	commonNamePanel := container.New(layout.NewVBoxLayout())
	languagesPanel := container.New(layout.NewVBoxLayout())
	flagPanel := container.New(layout.NewVBoxLayout())
	landlockedStatusPanel := container.New(layout.NewVBoxLayout())
	populationPanel := container.New(layout.NewVBoxLayout())
	currencyTypesPanel := container.New(layout.NewVBoxLayout())

	a := widget.NewLabel("Country's Official Name:")
	officialNamePanel.Add(a)
	officialNameDisplay = widget.NewEntry()
	officialNameDisplay.MultiLine = true
	officialNameDisplay.PlaceHolder = "Countries' official names will appear here"
	officialNameDisplay.Resize(fyne.NewSize(200, 100))
	officialNamePanel.Add(officialNameDisplay)
	detailsPanel.Add(officialNamePanel)

	a = widget.NewLabel("Country's Common Name:")
	commonNamePanel.Add(a)
	commonNameDisplay = widget.NewEntry()
	commonNameDisplay.MultiLine = true
	commonNameDisplay.PlaceHolder = "Countries' common names will appear here"
	commonNameDisplay.Resize(fyne.NewSize(200, 100))
	commonNamePanel.Add(commonNameDisplay)
	detailsPanel.Add(commonNamePanel)

	a = widget.NewLabel("Official Languages:")
	languagesPanel.Add(a)
	languagesDisplay = widget.NewEntry()
	languagesDisplay.MultiLine = true
	languagesDisplay.PlaceHolder = "Countries' official languages will appear here"
	languagesDisplay.Resize(fyne.NewSize(200, 100))
	languagesPanel.Add(languagesDisplay)
	detailsPanel.Add(languagesPanel)

	a = widget.NewLabel("National Flag")
	flagPanel.Add(a)
	//flagDisplay = widget.NewEntry() //remnants from flagDisplay being a widget like everything else
	//flagDisplay.MultiLine = true
	//flagDisplay.PlaceHolder = "Countries' flags will appear here"
	//flagDisplay.Resize(fyne.NewSize(200, 100))
	flagDisplay = &canvas.Image{}
	flagDisplay.FillMode = canvas.ImageFillContain
	flagDisplay.SetMinSize(fyne.NewSize(200, 100))

	flagPanel.Add(flagDisplay)
	detailsPanel.Add(flagPanel)

	a = widget.NewLabel("Country's Landlocked Status:")
	landlockedStatusPanel.Add(a)
	landlockedStatusDisplay = widget.NewEntry()
	landlockedStatusDisplay.MultiLine = true
	landlockedStatusDisplay.PlaceHolder = "Countries' landlocked status will appear here"
	landlockedStatusDisplay.Resize(fyne.NewSize(200, 100))
	landlockedStatusPanel.Add(landlockedStatusDisplay)
	detailsPanel.Add(landlockedStatusPanel)

	a = widget.NewLabel("Country's Population:")
	populationPanel.Add(a)
	populationDisplay = widget.NewEntry()
	populationDisplay.MultiLine = true
	populationDisplay.PlaceHolder = "Countries' population sizes will appear here"
	populationDisplay.Resize(fyne.NewSize(200, 100))
	populationPanel.Add(populationDisplay)
	detailsPanel.Add(populationPanel)

	a = widget.NewLabel("National Currencies:")
	currencyTypesPanel.Add(a)
	currencyTypesDisplay = widget.NewEntry()
	currencyTypesDisplay.MultiLine = true
	currencyTypesDisplay.PlaceHolder = "Countries' official currencies will appear here"
	currencyTypesDisplay.Resize(fyne.NewSize(200, 100))
	currencyTypesPanel.Add(currencyTypesDisplay)
	detailsPanel.Add(currencyTypesPanel)

	a = widget.NewLabel("Region:")
	countryRegionPanel.Add(a)
	countryRegionDisplay = widget.NewEntry()
	countryRegionDisplay.MultiLine = true
	countryRegionDisplay.PlaceHolder = "Countries' regions will appear here"
	countryRegionDisplay.Resize(fyne.NewSize(200, 100))
	countryRegionPanel.Add(countryRegionDisplay)
	detailsPanel.Add(countryRegionPanel)

	return detailsPanel
}
