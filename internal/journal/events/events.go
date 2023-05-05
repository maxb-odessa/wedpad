package events

import "github.com/maxb-odessa/slog"

type eventHandlersFunc func(map[string]interface{})

var eventHandlers map[string]eventHandlersFunc

func init() {
	eventHandlers = make(map[string]eventHandlersFunc)

	eventHandlers["AfmuRepairs"] = AfmuRepairs
	eventHandlers["ApproachBody"] = ApproachBody
	eventHandlers["ApproachSettlement"] = ApproachSettlement
	eventHandlers["Backpack"] = Backpack
	eventHandlers["BookTaxi"] = BookTaxi
	eventHandlers["BuyAmmo"] = BuyAmmo
	eventHandlers["BuyMicroResources"] = BuyMicroResources
	eventHandlers["BuySuit"] = BuySuit
	eventHandlers["BuyWeapon"] = BuyWeapon
	eventHandlers["Cargo"] = Cargo
	eventHandlers["CargoTransfer"] = CargoTransfer
	eventHandlers["CarrierBankTransfer"] = CarrierBankTransfer
	eventHandlers["CarrierBuy"] = CarrierBuy
	eventHandlers["CarrierCrewServices"] = CarrierCrewServices
	eventHandlers["CarrierDepositFuel"] = CarrierDepositFuel
	eventHandlers["CarrierFinance"] = CarrierFinance
	eventHandlers["CarrierJumpCancelled"] = CarrierJumpCancelled
	eventHandlers["CarrierJumpRequest"] = CarrierJumpRequest
	eventHandlers["CarrierStats"] = CarrierStats
	eventHandlers["CarrierTradeOrder"] = CarrierTradeOrder
	eventHandlers["CodexEntry"] = CodexEntry
	eventHandlers["CollectCargo"] = CollectCargo
	eventHandlers["Commander"] = Commander
	eventHandlers["CreateSuitLoadout"] = CreateSuitLoadout
	eventHandlers["DatalinkScan"] = DatalinkScan
	eventHandlers["DataScanned"] = DataScanned
	eventHandlers["DeleteSuitLoadout"] = DeleteSuitLoadout
	eventHandlers["DiscoveryScan"] = DiscoveryScan
	eventHandlers["Disembark"] = Disembark
	eventHandlers["Docked"] = Docked
	eventHandlers["DockingDenied"] = DockingDenied
	eventHandlers["DockingGranted"] = DockingGranted
	eventHandlers["DockingRequested"] = DockingRequested
	eventHandlers["DockSRV"] = DockSRV
	eventHandlers["EjectCargo"] = EjectCargo
	eventHandlers["Embark"] = Embark
	eventHandlers["EngineerCraft"] = EngineerCraft
	eventHandlers["EngineerProgress"] = EngineerProgress
	eventHandlers["FCMaterials"] = FCMaterials
	eventHandlers["FetchRemoteModule"] = FetchRemoteModule
	eventHandlers["Fileheader"] = Fileheader
	eventHandlers["Friends"] = Friends
	eventHandlers["FSDJump"] = FSDJump
	eventHandlers["FSDTarget"] = FSDTarget
	eventHandlers["FSSAllBodiesFound"] = FSSAllBodiesFound
	eventHandlers["FSSBodySignals"] = FSSBodySignals
	eventHandlers["FSSDiscoveryScan"] = FSSDiscoveryScan
	eventHandlers["FSSSignalDiscovered"] = FSSSignalDiscovered
	eventHandlers["FuelScoop"] = FuelScoop
	eventHandlers["HeatWarning"] = HeatWarning
	eventHandlers["HullDamage"] = HullDamage
	eventHandlers["Interdicted"] = Interdicted
	eventHandlers["JetConeBoost"] = JetConeBoost
	eventHandlers["LaunchSRV"] = LaunchSRV
	eventHandlers["LeaveBody"] = LeaveBody
	eventHandlers["Liftoff"] = Liftoff
	eventHandlers["LoadGame"] = LoadGame
	eventHandlers["Loadout"] = Loadout
	eventHandlers["Location"] = Location
	eventHandlers["Market"] = Market
	eventHandlers["MarketBuy"] = MarketBuy
	eventHandlers["MarketSell"] = MarketSell
	eventHandlers["MaterialCollected"] = MaterialCollected
	eventHandlers["MaterialDiscovered"] = MaterialDiscovered
	eventHandlers["Materials"] = Materials
	eventHandlers["MaterialTrade"] = MaterialTrade
	eventHandlers["MissionAccepted"] = MissionAccepted
	eventHandlers["MissionCompleted"] = MissionCompleted
	eventHandlers["MissionRedirected"] = MissionRedirected
	eventHandlers["Missions"] = Missions
	eventHandlers["ModuleBuy"] = ModuleBuy
	eventHandlers["ModuleBuyAndStore"] = ModuleBuyAndStore
	eventHandlers["ModuleInfo"] = ModuleInfo
	eventHandlers["ModuleRetrieve"] = ModuleRetrieve
	eventHandlers["ModuleSell"] = ModuleSell
	eventHandlers["ModuleSellRemote"] = ModuleSellRemote
	eventHandlers["ModuleStore"] = ModuleStore
	eventHandlers["MultiSellExplorationData"] = MultiSellExplorationData
	eventHandlers["Music"] = Music
	eventHandlers["NavRoute"] = NavRoute
	eventHandlers["NavRouteClear"] = NavRouteClear
	eventHandlers["null"] = null
	eventHandlers["Outfitting"] = Outfitting
	eventHandlers["Progress"] = Progress
	eventHandlers["Promotion"] = Promotion
	eventHandlers["Rank"] = Rank
	eventHandlers["ReceiveText"] = ReceiveText
	eventHandlers["RedeemVoucher"] = RedeemVoucher
	eventHandlers["RefuelAll"] = RefuelAll
	eventHandlers["Repair"] = Repair
	eventHandlers["Reputation"] = Reputation
	eventHandlers["ReservoirReplenished"] = ReservoirReplenished
	eventHandlers["RestockVehicle"] = RestockVehicle
	eventHandlers["Resurrect"] = Resurrect
	eventHandlers["SAAScanComplete"] = SAAScanComplete
	eventHandlers["SAASignalsFound"] = SAASignalsFound
	eventHandlers["Scan"] = Scan
	eventHandlers["ScanBaryCentre"] = ScanBaryCentre
	eventHandlers["Scanned"] = Scanned
	eventHandlers["ScanOrganic"] = ScanOrganic
	eventHandlers["Screenshot"] = Screenshot
	eventHandlers["SellOrganicData"] = SellOrganicData
	eventHandlers["SendText"] = SendText
	eventHandlers["SetUserShipName"] = SetUserShipName
	eventHandlers["ShieldState"] = ShieldState
	eventHandlers["ShipLocker"] = ShipLocker
	eventHandlers["ShipTargeted"] = ShipTargeted
	eventHandlers["Shipyard"] = Shipyard
	eventHandlers["ShipyardBuy"] = ShipyardBuy
	eventHandlers["ShipyardNew"] = ShipyardNew
	eventHandlers["ShipyardSell"] = ShipyardSell
	eventHandlers["ShipyardSwap"] = ShipyardSwap
	eventHandlers["ShipyardTransfer"] = ShipyardTransfer
	eventHandlers["Shutdown"] = Shutdown
	eventHandlers["SquadronStartup"] = SquadronStartup
	eventHandlers["StartJump"] = StartJump
	eventHandlers["Statistics"] = Statistics
	eventHandlers["Status"] = Status
	eventHandlers["StoredModules"] = StoredModules
	eventHandlers["StoredShips"] = StoredShips
	eventHandlers["SuitLoadout"] = SuitLoadout
	eventHandlers["SupercruiseEntry"] = SupercruiseEntry
	eventHandlers["SupercruiseExit"] = SupercruiseExit
	eventHandlers["SwitchSuitLoadout"] = SwitchSuitLoadout
	eventHandlers["Synthesis"] = Synthesis
	eventHandlers["TechnologyBroker"] = TechnologyBroker
	eventHandlers["Touchdown"] = Touchdown
	eventHandlers["UnderAttack"] = UnderAttack
	eventHandlers["Undocked"] = Undocked
	eventHandlers["USSDrop"] = USSDrop
}

func Handle(event string, data map[string]interface{}) {
	if handler, ok := eventHandlers[event]; ok {
		slog.Debug(9, "handling event '%s'", event)
		handler(data)
	} else {
		slog.Err("Undefined event '%s', can't handle", event)
	}
}
