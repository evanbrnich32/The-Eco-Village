package village

import "fmt"

// TheEcoVillage is a type of village that is environmentally friendly and sustainable
type TheEcoVillage struct {
	Name        string
	Description string
}

// NewTheEcoVillage creates a new TheEcoVillage
func NewTheEcoVillage(name string, description string) *TheEcoVillage {
	return &TheEcoVillage{
		Name:        name,
		Description: description,
	}
}

// GetDescription prints the description of the village
func (v *TheEcoVillage) GetDescription() {
	fmt.Printf("The Eco Village is %s \n", v.Description)
}

// GetName returns the name of the village
func (v *TheEcoVillage) GetName() string {
	return v.Name
}

// SetName changes the name of the village
func (v *TheEcoVillage) SetName(name string) {
	v.Name = name
}

// SetDescription changes the description of the village
func (v *TheEcoVillage) SetDescription(description string) {
	v.Description = description
}

// PrintDetails prints out the details of the village
func (v *TheEcoVillage) PrintDetails() {
	fmt.Printf("The Eco Village is called %s and is %s \n", v.Name, v.Description)
}

// CreateSustainableEnergy creates sustainable energy for the village
func (v *TheEcoVillage) CreateSustainableEnergy() {
	fmt.Println("The Eco Village is using solar power and wind turbines to create sustainable energy")
}

// PlantTrees plants enough trees to offset the carbon footprint of the village
func (v *TheEcoVillage) PlantTrees() {
	fmt.Println("The Eco Village is planting enough trees to offset its carbon footprint")
}

// UseOrganicMaterials uses organic materials to build buildings in the village
func (v *TheEcoVillage) UseOrganicMaterials() {
	fmt.Println("The Eco Village is using organic materials to build its buildings")
}

// MakeEcoFriendlyJobs creates eco-friendly jobs in the village
func (v *TheEcoVillage) MakeEcoFriendlyJobs() {
	fmt.Println("The Eco Village is creating eco-friendly jobs such as farming and gardening")
}

// UseRainwaterHarvesting implements rainwater harvesting to collect water
func (v *TheEcoVillage) UseRainwaterHarvesting() {
	fmt.Println("The Eco Village is using rainwater harvesting to collect water")
}

// InstallRecyclingSystem installs a recycling system in the village
func (v *TheEcoVillage) InstallRecyclingSystem() {
	fmt.Println("The Eco Village is installing a recycling system to reduce waste")
}

// UseRenewableEnergy sources renewable energy to power the village
func (v *TheEcoVillage) UseRenewableEnergy() {
	fmt.Println("The Eco Village is using renewable energy to power its needs")
}

// IncreaseGreenSpace increases the available green space in the village
func (v *TheEcoVillage) IncreaseGreenSpace() {
	fmt.Println("The Eco Village is increasing its green space for recreational use")
}

// DevelopGreenTransportation implements green transportation to the village
func (v *TheEcoVillage) DevelopGreenTransportation() {
	fmt.Println("The Eco Village is developing green transportation for its citizens")
}

// ReduceExternalNoise implements methods to reduce external noise
func (v *TheEcoVillage) ReduceExternalNoise() {
	fmt.Println("The Eco Village is reducing external noise with Sound Barrier Walls")
}

// UtilizeBiofuel uses biofuel to reduce emissions
func (v *TheEcoVillage) UtilizeBiofuel() {
	fmt.Println("The Eco Village is utilizing biofuel to reduce emissions")
}

// PromoteOrganicFarming promotes organic farming to reduce the use of chemicals
func (v *TheEcoVillage) PromoteOrganicFarming() {
	fmt.Println("The Eco Village is promoting organic farming to reduce the use of chemicals")
}

// ConserveWater implements water conservation techniques to reduce water use
func (v *TheEcoVillage) ConserveWater() {
	fmt.Println("The Eco Village is implementing water conservation techniques to reduce water use")
}

// InstallSolarPanels install solar panels to generate renewable energy
func (v *TheEcoVillage) InstallSolarPanels() {
	fmt.Println("The Eco Village is installing solar panels to generate renewable energy")
}

// ImplementComposting implements composting to reduce waste
func (v *TheEcoVillage) ImplementComposting() {
	fmt.Println("The Eco Village is implementing composting to reduce waste")
}

// ConserveBiodiversity includes conservation efforts to protect the local biodiversity
func (v *TheEcoVillage) ConserveBiodiversity() {
	fmt.Println("The Eco Village is conserving biodiversity to protect the local ecosystem")
}

// UseEnvironmentallyFriendlyMaterials uses environment-friendly materials while constructing buildings
func (v *TheEcoVillage) UseEnvironmentallyFriendlyMaterials() {
	fmt.Println("The Eco Village is using eco-friendly materials while constructing buildings")
}

// RegulateResourceUse implements regulations on water and energy use
func (v *TheEcoVillage) RegulateResourceUse() {
	fmt.Println("The Eco Village is implementing regulations on water and energy use")
}

// UtilizeGreenArchitecture implements green architecture in the village
func (v *TheEcoVillage) UtilizeGreenArchitecture() {
	fmt.Println("The Eco Village is utilizing green architecture to reduce its environmental impact")
}

// PromoteCommunityGardening encourages community gardening to boost local food production
func (v *TheEcoVillage) PromoteCommunityGardening() {
	fmt.Println("The Eco Village is promoting community gardening to boost local food production")
}

// UseIntegratedWaterManagement implements integrated water management practices
func (v *TheEcoVillage) UseIntegratedWaterManagement() {
	fmt.Println("The Eco Village is using integrated water management practices to conserve water")
}

// PromoteRecycling encourages recycling
func (v *TheEcoVillage) PromoteRecycling() {
	fmt.Println("The Eco Village is promoting recycling to reduce waste")
}

// DevelopGreenTechnologies develops green technologies for the village
func (v *TheEcoVillage) DevelopGreenTechnologies() {
	fmt.Println("The Eco Village is developing green technologies to minimize its environmental impact")
}