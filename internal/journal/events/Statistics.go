package events

import (
	"time"
)

// Statistics event structure
type StatisticsT struct {
	BankAccount struct {
		CurrentWealth          int `mapstructure:"Current_Wealth"`
		InsuranceClaims        int `mapstructure:"Insurance_Claims"`
		OwnedShipCount         int `mapstructure:"Owned_Ship_Count"`
		PremiumStockBought     int `mapstructure:"Premium_Stock_Bought"`
		SpentOnAmmoConsumables int `mapstructure:"Spent_On_Ammo_Consumables"`
		SpentOnFuel            int `mapstructure:"Spent_On_Fuel"`
		SpentOnInsurance       int `mapstructure:"Spent_On_Insurance"`
		SpentOnOutfitting      int `mapstructure:"Spent_On_Outfitting"`
		SpentOnPremiumStock    int `mapstructure:"Spent_On_Premium_Stock"`
		SpentOnRepairs         int `mapstructure:"Spent_On_Repairs"`
		SpentOnShips           int `mapstructure:"Spent_On_Ships"`
		SpentOnSuitConsumables int `mapstructure:"Spent_On_Suit_Consumables"`
		SpentOnSuits           int `mapstructure:"Spent_On_Suits"`
		SpentOnWeapons         int `mapstructure:"Spent_On_Weapons"`
		SuitsOwned             int `mapstructure:"Suits_Owned"`
		WeaponsOwned           int `mapstructure:"Weapons_Owned"`
	} `mapstructure:"Bank_Account"`
	Combat struct {
		AssassinationProfits     int `mapstructure:"Assassination_Profits"`
		Assassinations           int `mapstructure:"Assassinations"`
		BountiesClaimed          int `mapstructure:"Bounties_Claimed"`
		BountyHuntingProfit      int `mapstructure:"Bounty_Hunting_Profit"`
		CombatBondProfits        int `mapstructure:"Combat_Bond_Profits"`
		CombatBonds              int `mapstructure:"Combat_Bonds"`
		ConflictZoneHigh         int `mapstructure:"ConflictZone_High"`
		ConflictZoneHighWins     int `mapstructure:"ConflictZone_High_Wins"`
		ConflictZoneLow          int `mapstructure:"ConflictZone_Low"`
		ConflictZoneLowWins      int `mapstructure:"ConflictZone_Low_Wins"`
		ConflictZoneMedium       int `mapstructure:"ConflictZone_Medium"`
		ConflictZoneMediumWins   int `mapstructure:"ConflictZone_Medium_Wins"`
		ConflictZoneTotal        int `mapstructure:"ConflictZone_Total"`
		ConflictZoneTotalWins    int `mapstructure:"ConflictZone_Total_Wins"`
		DropshipsBooked          int `mapstructure:"Dropships_Booked"`
		DropshipsCancelled       int `mapstructure:"Dropships_Cancelled"`
		DropshipsTaken           int `mapstructure:"Dropships_Taken"`
		HighestSingleReward      int `mapstructure:"Highest_Single_Reward"`
		OnFootCombatBonds        int `mapstructure:"OnFoot_Combat_Bonds"`
		OnFootCombatBondsProfits int `mapstructure:"OnFoot_Combat_Bonds_Profits"`
		OnFootScavsKilled        int `mapstructure:"OnFoot_Scavs_Killed"`
		OnFootShipsDestroyed     int `mapstructure:"OnFoot_Ships_Destroyed"`
		OnFootSkimmersKilled     int `mapstructure:"OnFoot_Skimmers_Killed"`
		OnFootVehiclesDestroyed  int `mapstructure:"OnFoot_Vehicles_Destroyed"`
		SettlementConquered      int `mapstructure:"Settlement_Conquered"`
		SettlementDefended       int `mapstructure:"Settlement_Defended"`
		SkimmersKilled           int `mapstructure:"Skimmers_Killed"`
	} `mapstructure:"Combat"`
	Crafting struct {
		CountOfUsedEngineers  int `mapstructure:"Count_Of_Used_Engineers"`
		RecipesGenerated      int `mapstructure:"Recipes_Generated"`
		RecipesGeneratedRank1 int `mapstructure:"Recipes_Generated_Rank_1"`
		RecipesGeneratedRank2 int `mapstructure:"Recipes_Generated_Rank_2"`
		RecipesGeneratedRank3 int `mapstructure:"Recipes_Generated_Rank_3"`
		RecipesGeneratedRank4 int `mapstructure:"Recipes_Generated_Rank_4"`
		RecipesGeneratedRank5 int `mapstructure:"Recipes_Generated_Rank_5"`
		SuitModsApplied       int `mapstructure:"Suit_Mods_Applied"`
		SuitModsAppliedFull   int `mapstructure:"Suit_Mods_Applied_Full"`
		SuitsUpgraded         int `mapstructure:"Suits_Upgraded"`
		SuitsUpgradedFull     int `mapstructure:"Suits_Upgraded_Full"`
		WeaponModsApplied     int `mapstructure:"Weapon_Mods_Applied"`
		WeaponModsAppliedFull int `mapstructure:"Weapon_Mods_Applied_Full"`
		WeaponsUpgraded       int `mapstructure:"Weapons_Upgraded"`
		WeaponsUpgradedFull   int `mapstructure:"Weapons_Upgraded_Full"`
	} `mapstructure:"Crafting"`
	Crew struct {
		NpcCrewDied       int `mapstructure:"NpcCrew_Died"`
		NpcCrewFired      int `mapstructure:"NpcCrew_Fired"`
		NpcCrewHired      int `mapstructure:"NpcCrew_Hired"`
		NpcCrewTotalWages int `mapstructure:"NpcCrew_TotalWages"`
	} `mapstructure:"Crew"`
	Crime struct {
		BountiesReceived         int `mapstructure:"Bounties_Received"`
		CitizensMurdered         int `mapstructure:"Citizens_Murdered"`
		DataStolen               int `mapstructure:"Data_Stolen"`
		Fines                    int `mapstructure:"Fines"`
		GoodsStolen              int `mapstructure:"Goods_Stolen"`
		GuardsMurdered           int `mapstructure:"Guards_Murdered"`
		HighestBounty            int `mapstructure:"Highest_Bounty"`
		MalwareUploaded          int `mapstructure:"Malware_Uploaded"`
		Notoriety                int `mapstructure:"Notoriety"`
		OmnipolMurdered          int `mapstructure:"Omnipol_Murdered"`
		ProductionSabotage       int `mapstructure:"Production_Sabotage"`
		ProductionTheft          int `mapstructure:"Production_Theft"`
		ProfilesCloned           int `mapstructure:"Profiles_Cloned"`
		SampleStolen             int `mapstructure:"Sample_Stolen"`
		SettlementsStateShutdown int `mapstructure:"Settlements_State_Shutdown"`
		TotalBounties            int `mapstructure:"Total_Bounties"`
		TotalFines               int `mapstructure:"Total_Fines"`
		TotalMurders             int `mapstructure:"Total_Murders"`
		TotalStolen              int `mapstructure:"Total_Stolen"`
		TurretsDestroyed         int `mapstructure:"Turrets_Destroyed"`
		TurretsOverloaded        int `mapstructure:"Turrets_Overloaded"`
		TurretsTotal             int `mapstructure:"Turrets_Total"`
		ValueStolenStateChange   int `mapstructure:"Value_Stolen_StateChange"`
	} `mapstructure:"Crime"`
	Exobiology struct {
		FirstLogged               int `mapstructure:"First_Logged"`
		FirstLoggedProfits        int `mapstructure:"First_Logged_Profits"`
		OrganicData               int `mapstructure:"Organic_Data"`
		OrganicDataProfits        int `mapstructure:"Organic_Data_Profits"`
		OrganicGenus              int `mapstructure:"Organic_Genus"`
		OrganicGenusEncountered   int `mapstructure:"Organic_Genus_Encountered"`
		OrganicPlanets            int `mapstructure:"Organic_Planets"`
		OrganicSpecies            int `mapstructure:"Organic_Species"`
		OrganicSpeciesEncountered int `mapstructure:"Organic_Species_Encountered"`
		OrganicSystems            int `mapstructure:"Organic_Systems"`
		OrganicVariantEncountered int `mapstructure:"Organic_Variant_Encountered"`
	} `mapstructure:"Exobiology"`
	Exploration struct {
		EfficientScans            int     `mapstructure:"Efficient_Scans"`
		ExplorationProfits        int     `mapstructure:"Exploration_Profits"`
		FirstFootfalls            int     `mapstructure:"First_Footfalls"`
		GreatestDistanceFromStart float64 `mapstructure:"Greatest_Distance_From_Start"`
		HighestPayout             int     `mapstructure:"Highest_Payout"`
		OnFootDistanceTravelled   int     `mapstructure:"OnFoot_Distance_Travelled"`
		PlanetFootfalls           int     `mapstructure:"Planet_Footfalls"`
		PlanetsScannedToLevel2    int     `mapstructure:"Planets_Scanned_To_Level_2"`
		PlanetsScannedToLevel3    int     `mapstructure:"Planets_Scanned_To_Level_3"`
		SettlementsVisited        int     `mapstructure:"Settlements_Visited"`
		ShuttleDistanceTravelled  float64 `mapstructure:"Shuttle_Distance_Travelled"`
		ShuttleJourneys           int     `mapstructure:"Shuttle_Journeys"`
		SpentOnShuttles           int     `mapstructure:"Spent_On_Shuttles"`
		SystemsVisited            int     `mapstructure:"Systems_Visited"`
		TimePlayed                int     `mapstructure:"Time_Played"`
		TotalHyperspaceDistance   int     `mapstructure:"Total_Hyperspace_Distance"`
		TotalHyperspaceJumps      int     `mapstructure:"Total_Hyperspace_Jumps"`
	} `mapstructure:"Exploration"`
	Fleetcarrier *struct {
		FleetcarrierDistanceTravelled float64 `mapstructure:"FLEETCARRIER_DISTANCE_TRAVELLED"`
		FleetcarrierExportTotal       int     `mapstructure:"FLEETCARRIER_EXPORT_TOTAL"`
		FleetcarrierImportTotal       int     `mapstructure:"FLEETCARRIER_IMPORT_TOTAL"`
		FleetcarrierOutfittingProfit  int     `mapstructure:"FLEETCARRIER_OUTFITTING_PROFIT"`
		FleetcarrierOutfittingSold    int     `mapstructure:"FLEETCARRIER_OUTFITTING_SOLD"`
		FleetcarrierRearmTotal        int     `mapstructure:"FLEETCARRIER_REARM_TOTAL"`
		FleetcarrierRefuelProfit      int     `mapstructure:"FLEETCARRIER_REFUEL_PROFIT"`
		FleetcarrierRefuelTotal       int     `mapstructure:"FLEETCARRIER_REFUEL_TOTAL"`
		FleetcarrierRepairsTotal      int     `mapstructure:"FLEETCARRIER_REPAIRS_TOTAL"`
		FleetcarrierShipyardProfit    int     `mapstructure:"FLEETCARRIER_SHIPYARD_PROFIT"`
		FleetcarrierShipyardSold      int     `mapstructure:"FLEETCARRIER_SHIPYARD_SOLD"`
		FleetcarrierStolenprofitTotal int     `mapstructure:"FLEETCARRIER_STOLENPROFIT_TOTAL"`
		FleetcarrierStolenspendTotal  int     `mapstructure:"FLEETCARRIER_STOLENSPEND_TOTAL"`
		FleetcarrierTotalJumps        int     `mapstructure:"FLEETCARRIER_TOTAL_JUMPS"`
		FleetcarrierTradeprofitTotal  int     `mapstructure:"FLEETCARRIER_TRADEPROFIT_TOTAL"`
		FleetcarrierTradespendTotal   int     `mapstructure:"FLEETCARRIER_TRADESPEND_TOTAL"`
		FleetcarrierVouchersProfit    int     `mapstructure:"FLEETCARRIER_VOUCHERS_PROFIT"`
		FleetcarrierVouchersRedeemed  int     `mapstructure:"FLEETCARRIER_VOUCHERS_REDEEMED"`
	} `mapstructure:"FLEETCARRIER"`
	MaterialTraderStats struct {
		AssetsTradedIn         int `mapstructure:"Assets_Traded_In"`
		AssetsTradedOut        int `mapstructure:"Assets_Traded_Out"`
		EncodedMaterialsTraded int `mapstructure:"Encoded_Materials_Traded"`
		Grade1MaterialsTraded  int `mapstructure:"Grade_1_Materials_Traded"`
		Grade2MaterialsTraded  int `mapstructure:"Grade_2_Materials_Traded"`
		Grade3MaterialsTraded  int `mapstructure:"Grade_3_Materials_Traded"`
		Grade4MaterialsTraded  int `mapstructure:"Grade_4_Materials_Traded"`
		Grade5MaterialsTraded  int `mapstructure:"Grade_5_Materials_Traded"`
		MaterialsTraded        int `mapstructure:"Materials_Traded"`
		RawMaterialsTraded     int `mapstructure:"Raw_Materials_Traded"`
		TradesCompleted        int `mapstructure:"Trades_Completed"`
	} `mapstructure:"Material_Trader_Stats"`
	Mining struct {
		MaterialsCollected int `mapstructure:"Materials_Collected"`
		MiningProfits      int `mapstructure:"Mining_Profits"`
		QuantityMined      int `mapstructure:"Quantity_Mined"`
	} `mapstructure:"Mining"`
	Multicrew struct {
		MulticrewCreditsTotal     int `mapstructure:"Multicrew_Credits_Total"`
		MulticrewFighterTimeTotal int `mapstructure:"Multicrew_Fighter_Time_Total"`
		MulticrewFinesTotal       int `mapstructure:"Multicrew_Fines_Total"`
		MulticrewGunnerTimeTotal  int `mapstructure:"Multicrew_Gunner_Time_Total"`
		MulticrewTimeTotal        int `mapstructure:"Multicrew_Time_Total"`
	} `mapstructure:"Multicrew"`
	Passengers struct {
		PassengersMissionsAccepted    int `mapstructure:"Passengers_Missions_Accepted"`
		PassengersMissionsBulk        int `mapstructure:"Passengers_Missions_Bulk"`
		PassengersMissionsDelivered   int `mapstructure:"Passengers_Missions_Delivered"`
		PassengersMissionsDisgruntled int `mapstructure:"Passengers_Missions_Disgruntled"`
		PassengersMissionsEjected     int `mapstructure:"Passengers_Missions_Ejected"`
		PassengersMissionsVip         int `mapstructure:"Passengers_Missions_VIP"`
	} `mapstructure:"Passengers"`
	SearchAndRescue struct {
		MaglocksOpened            int `mapstructure:"Maglocks_Opened"`
		PanelsOpened              int `mapstructure:"Panels_Opened"`
		SalvageIllegalPoi         int `mapstructure:"Salvage_Illegal_POI"`
		SalvageIllegalSettlements int `mapstructure:"Salvage_Illegal_Settlements"`
		SalvageLegalPoi           int `mapstructure:"Salvage_Legal_POI"`
		SalvageLegalSettlements   int `mapstructure:"Salvage_Legal_Settlements"`
		SearchRescueCount         int `mapstructure:"SearchRescue_Count"`
		SearchRescueProfit        int `mapstructure:"SearchRescue_Profit"`
		SearchRescueTraded        int `mapstructure:"SearchRescue_Traded"`
		SettlementsStateFireOut   int `mapstructure:"Settlements_State_FireOut"`
		SettlementsStateReboot    int `mapstructure:"Settlements_State_Reboot"`
	} `mapstructure:"Search_And_Rescue"`
	Smuggling struct {
		AverageProfit            int `mapstructure:"Average_Profit"`
		BlackMarketsProfits      int `mapstructure:"Black_Markets_Profits"`
		BlackMarketsTradedWith   int `mapstructure:"Black_Markets_Traded_With"`
		HighestSingleTransaction int `mapstructure:"Highest_Single_Transaction"`
		ResourcesSmuggled        int `mapstructure:"Resources_Smuggled"`
	} `mapstructure:"Smuggling"`
	Trading struct {
		AssetsSold               int     `mapstructure:"Assets_Sold"`
		AverageProfit            float64 `mapstructure:"Average_Profit"`
		DataSold                 int     `mapstructure:"Data_Sold"`
		GoodsSold                int     `mapstructure:"Goods_Sold"`
		HighestSingleTransaction int     `mapstructure:"Highest_Single_Transaction"`
		MarketProfits            int     `mapstructure:"Market_Profits"`
		MarketsTradedWith        int     `mapstructure:"Markets_Traded_With"`
		ResourcesTraded          int     `mapstructure:"Resources_Traded"`
	} `mapstructure:"Trading"`
	Event     string    `mapstructure:"event"`
	Timestamp time.Time `mapstructure:"timestamp"`
}

// Statistics event handler
func (evHandler EventHandler) Statistics(eventData map[string]interface{}) {
    // ev := new(StatisticsT)
    // mapstructure.Decode(eventData, ev)
}

