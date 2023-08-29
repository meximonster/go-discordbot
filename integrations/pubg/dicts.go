package pubg

var maps = map[string]string{
	"Baltic_Main":     "Erangel (Remastered)",
	"Chimera_Main":    "Paramo",
	"Desert_Main":     "Miramar",
	"DihorOtok_Main":  "Vikendi",
	"Erangel_Main":    "Erangel",
	"Heaven_Main":     "Haven",
	"Kiki_Main":       "Deston",
	"Range_Main":      "Camp Jackal",
	"Savage_Main":     "Sanhok",
	"Summerland_Main": "Karakin",
	"Tiger_Main":      "Taego",
}

var damageCauser = map[string]string{
	"AIPawn_Base_Female_C":                    "AI Player",
	"AIPawn_Base_Male_C":                      "AI Player",
	"AirBoat_V2_C":                            "Airboat",
	"AquaRail_A_01_C":                         "Aquarail",
	"AquaRail_A_02_C":                         "Aquarail",
	"AquaRail_A_03_C":                         "Aquarail",
	"BP_ATV_C":                                "Quad",
	"BP_BearV2_C":                             "Bear",
	"BP_BRDM_C":                               "BRDM-2",
	"BP_Bicycle_C":                            "Mountain Bike",
	"BP_CoupeRB_C":                            "Coupe RB",
	"BP_DO_Circle_Train_Merged_C":             "Train",
	"BP_DO_Line_Train_Dino_Merged_C":          "Train",
	"BP_DO_Line_Train_Merged_C":               "Train",
	"BP_Dirtbike_C":                           "Dirt Bike",
	"BP_DronePackage_Projectile_C":            "Drone",
	"BP_Eragel_CargoShip01_C":                 "Ferry Damage",
	"BP_FakeLootProj_AmmoBox_C":               "Loot Truck",
	"BP_FakeLootProj_MilitaryCrate_C":         "Loot Truck",
	"BP_FireEffectController_C":               "Molotov Fire",
	"BP_FireEffectController_JerryCan_C":      "Jerrycan Fire",
	"BP_Food_Truck_C":                         "Food Truck",
	"BP_Helicopter_C":                         "Pillar Scout Helicopter",
	"BP_IncendiaryDebuff_C":                   "Burn",
	"BP_JerryCanFireDebuff_C":                 "Burn",
	"BP_JerryCan_FuelPuddle_C":                "Burn",
	"BP_KillTruck_C":                          "Kill Truck",
	"BP_LootTruck_C":                          "Loot Truck",
	"BP_M_Rony_A_01_C":                        "Rony",
	"BP_M_Rony_A_02_C":                        "Rony",
	"BP_M_Rony_A_03_C":                        "Rony",
	"BP_Mirado_A_02_C":                        "Mirado",
	"BP_Mirado_A_03_Esports_C":                "Mirado",
	"BP_Mirado_Open_03_C":                     "Mirado (open top)",
	"BP_Mirado_Open_04_C":                     "Mirado (open top)",
	"BP_Mirado_Open_05_C":                     "Mirado (open top)",
	"BP_MolotovFireDebuff_C":                  "Molotov Fire Damage",
	"BP_Motorbike_04_C":                       "Motorcycle",
	"BP_Motorbike_04_Desert_C":                "Motorcycle",
	"BP_Motorbike_04_SideCar_C":               "Motorcycle (w/ Sidecar)",
	"BP_Motorbike_04_SideCar_Desert_C":        "Motorcycle (w/ Sidecar)",
	"BP_Motorbike_Solitario_C":                "Motorcycle",
	"BP_Motorglider_C":                        "Motor Glider",
	"BP_Motorglider_Green_C":                  "Motor Glider",
	"BP_Niva_01_C":                            "Zima",
	"BP_Niva_02_C":                            "Zima",
	"BP_Niva_03_C":                            "Zima",
	"BP_Niva_04_C":                            "Zima",
	"BP_Niva_05_C":                            "Zima",
	"BP_Niva_06_C":                            "Zima",
	"BP_Niva_07_C":                            "Zima",
	"BP_PickupTruck_A_01_C":                   "Pickup Truck (closed top)",
	"BP_PickupTruck_A_02_C":                   "Pickup Truck (closed top)",
	"BP_PickupTruck_A_03_C":                   "Pickup Truck (closed top)",
	"BP_PickupTruck_A_04_C":                   "Pickup Truck (closed top)",
	"BP_PickupTruck_A_05_C":                   "Pickup Truck (closed top)",
	"BP_PickupTruck_A_esports_C":              "Pickup Truck (closed top)",
	"BP_PickupTruck_B_01_C":                   "Pickup Truck (open top)",
	"BP_PickupTruck_B_02_C":                   "Pickup Truck (open top)",
	"BP_PickupTruck_B_03_C":                   "Pickup Truck (open top)",
	"BP_PickupTruck_B_04_C":                   "Pickup Truck (open top)",
	"BP_PickupTruck_B_05_C":                   "Pickup Truck (open top)",
	"BP_Pillar_Car_C":                         "Pillar Security Car",
	"BP_PonyCoupe_C":                          "Pony Coupe",
	"BP_Porter_C":                             "Porter",
	"BP_Scooter_01_A_C":                       "Scooter",
	"BP_Scooter_02_A_C":                       "Scooter",
	"BP_Scooter_03_A_C":                       "Scooter",
	"BP_Scooter_04_A_C":                       "Scooter",
	"BP_Snowbike_01_C":                        "Snowbike",
	"BP_Snowbike_02_C":                        "Snowbike",
	"BP_Snowmobile_01_C":                      "Snowmobile",
	"BP_Snowmobile_02_C":                      "Snowmobile",
	"BP_Snowmobile_03_C":                      "Snowmobile",
	"BP_Spiketrap_C":                          "Spike Trap",
	"BP_TslGasPump_C":                         "Gas Pump",
	"BP_TukTukTuk_A_01_C":                     "Tukshai",
	"BP_TukTukTuk_A_02_C":                     "Tukshai",
	"BP_TukTukTuk_A_03_C":                     "Tukshai",
	"BP_Van_A_01_C":                           "Van",
	"BP_Van_A_02_C":                           "Van",
	"BP_Van_A_03_C":                           "Van",
	"BattleRoyaleModeController_Chimera_C":    "Bluezone",
	"BattleRoyaleModeController_Def_C":        "Bluezone",
	"BattleRoyaleModeController_Desert_C":     "Bluezone",
	"BattleRoyaleModeController_DihorOtok_C":  "Bluezone",
	"BattleRoyaleModeController_Heaven_C":     "Bluezone",
	"BattleRoyaleModeController_Kiki_C":       "Bluezone",
	"BattleRoyaleModeController_Savage_C":     "Bluezone",
	"BattleRoyaleModeController_Summerland_C": "Bluezone",
	"BattleRoyaleModeController_Tiger_C":      "Bluezone",
	"BlackZoneController_Def_C":               "Blackzone",
	"Bluezonebomb_EffectActor_C":              "Bluezone Grenade",
	"Boat_PG117_C":                            "PG-117",
	"Buff_DecreaseBreathInApnea_C":            "Drowning",
	"Buggy_A_01_C":                            "Buggy",
	"Buggy_A_02_C":                            "Buggy",
	"Buggy_A_03_C":                            "Buggy",
	"Buggy_A_04_C":                            "Buggy",
	"Buggy_A_05_C":                            "Buggy",
	"Buggy_A_06_C":                            "Buggy",
	"Carepackage_Container_C":                 "Care Package",
	"Dacia_A_01_v2_C":                         "Dacia",
	"Dacia_A_01_v2_snow_C":                    "Dacia",
	"Dacia_A_02_v2_C":                         "Dacia",
	"Dacia_A_03_v2_C":                         "Dacia",
	"Dacia_A_03_v2_Esports_C":                 "Dacia",
	"Dacia_A_04_v2_C":                         "Dacia",
	"DroppedItemGroup":                        "Object Fragments",
	"EmergencyAircraft_Tiger_C":               "Emergency Aircraft",
	"Jerrycan":                                "Jerrycan",
	"JerrycanFire":                            "Jerrycan Fire",
	"Lava":                                    "Lava",
	"Mortar_Projectile_C":                     "Mortar Projectile",
	"None":                                    "None",
	"PG117_A_01_C":                            "PG-117",
	"PanzerFaust100M_Projectile_C":            "Panzerfaust Projectile",
	"PlayerFemale_A_C":                        "Player",
	"PlayerMale_A_C":                          "Player",
	"ProjC4_C":                                "C4",
	"ProjGrenade_C":                           "Frag Grenade",
	"ProjIncendiary_C":                        "Incendiary Projectile",
	"ProjMolotov_C":                           "Molotov Cocktail",
	"ProjMolotov_DamageField_Direct_C":        "Molotov Cocktail Fire Field",
	"ProjStickyGrenade_C":                     "Sticky Bomb",
	"RacingDestructiblePropaneTankActor_C":    "Propane Tank",
	"RacingModeContorller_Desert_C":           "Bluezone",
	"RedZoneBomb_C":                           "Redzone",
	"RedZoneBombingField":                     "Redzone",
	"RedZoneBombingField_Def_C":               "Redzone",
	"TslDestructibleSurfaceManager":           "Destructible Surface",
	"TslPainCausingVolume":                    "Lava",
	"Uaz_A_01_C":                              "UAZ (open top)",
	"Uaz_Armored_C":                           "UAZ (armored)",
	"Uaz_B_01_C":                              "UAZ (soft top)",
	"Uaz_B_01_esports_C":                      "UAZ (soft top)",
	"Uaz_C_01_C":                              "UAZ (hard top)",
	"UltAIPawn_Base_Female_C":                 "Player",
	"UltAIPawn_Base_Male_C":                   "Player",
	"WeapACE32_C":                             "ACE32",
	"WeapAK47_C":                              "AKM",
	"WeapAUG_C":                               "AUG A3",
	"WeapAWM_C":                               "AWM",
	"WeapBerreta686_C":                        "S686",
	"WeapBerylM762_C":                         "Beryl",
	"WeapBizonPP19_C":                         "Bizon",
	"WeapCowbarProjectile_C":                  "Crowbar Projectile",
	"WeapCowbar_C":                            "Crowbar",
	"WeapCrossbow_1_C":                        "Crossbow",
	"WeapDP12_C":                              "DBS",
	"WeapDP28_C":                              "DP-28",
	"WeapDesertEagle_C":                       "Deagle",
	"WeapDuncansHK416_C":                      "M416",
	"WeapFNFal_C":                             "SLR",
	"WeapG18_C":                               "P18C",
	"WeapG36C_C":                              "G36C",
	"WeapGroza_C":                             "Groza",
	"WeapHK416_C":                             "M416",
	"WeapJuliesKar98k_C":                      "Kar98k",
	"WeapK2_C":                                "K2",
	"WeapKar98k_C":                            "Kar98k",
	"WeapL6_C":                                "Lynx AMR",
	"WeapLunchmeatsAK47_C":                    "AKM",
	"WeapM16A4_C":                             "M16A4",
	"WeapM1911_C":                             "P1911",
	"WeapM249_C":                              "M249",
	"WeapM24_C":                               "M24",
	"WeapM9_C":                                "P92",
	"WeapMG3_C":                               "MG3",
	"WeapMP5K_C":                              "MP5K",
	"WeapMP9_C":                               "MP9",
	"WeapMacheteProjectile_C":                 "Machete Projectile",
	"WeapMachete_C":                           "Machete",
	"WeapMadsQBU88_C":                         "QBU88",
	"WeapMini14_C":                            "Mini 14",
	"WeapMk12_C":                              "Mk12",
	"WeapMk14_C":                              "Mk14 EBR",
	"WeapMk47Mutant_C":                        "Mk47 Mutant",
	"WeapMosinNagant_C":                       "Mosin-Nagant",
	"WeapNagantM1895_C":                       "R1895",
	"WeapOriginS12_C":                         "O12",
	"WeapP90_C":                               "P90",
	"WeapPanProjectile_C":                     "Pan Projectile",
	"WeapPan_C":                               "Pan",
	"WeapPanzerFaust100M1_C":                  "Panzerfaust",
	"WeapQBU88_C":                             "QBU88",
	"WeapQBZ95_C":                             "QBZ95",
	"WeapRhino_C":                             "R45",
	"WeapSCAR-L_C":                            "SCAR-L",
	"WeapSKS_C":                               "SKS",
	"WeapSaiga12_C":                           "S12K",
	"WeapSawnoff_C":                           "Sawed-off",
	"WeapSickleProjectile_C":                  "Sickle Projectile",
	"WeapSickle_C":                            "Sickle",
	"WeapThompson_C":                          "Tommy Gun",
	"WeapTurret_KillTruck_Main_C":             "Kill Truck Turret",
	"WeapUMP_C":                               "UMP9",
	"WeapUZI_C":                               "Micro Uzi",
	"WeapVSS_C":                               "VSS",
	"WeapVector_C":                            "Vector",
	"WeapWin94_C":                             "Win94",
	"WeapWinchester_C":                        "S1897",
	"Weapvz61Skorpion_C":                      "Skorpion",
}

var items = map[string]string{
	"Helmet_Repair_Kit_C":                                         "Helmet Repair Kit",
	"InstantRevivalKit_C":                                         "Critical Response Kit",
	"Item_Ammo_12GuageSlug_C":                                     "12 Gauge Slug",
	"Item_Ammo_12Guage_C":                                         "12 Gauge Ammo",
	"Item_Ammo_300Magnum_C":                                       "300 Magnum Ammo",
	"Item_Ammo_40mm_C":                                            "40mm Smoke Grenade",
	"Item_Ammo_45ACP_C":                                           ".45 ACP Ammo",
	"Item_Ammo_556mm_C":                                           "5.56mm Ammo",
	"Item_Ammo_57mm_C":                                            "57mm Ammo",
	"Item_Ammo_762mm_C":                                           "7.62mm Ammo",
	"Item_Ammo_9mm_C":                                             "9mm Ammo",
	"Item_Ammo_Bolt_C":                                            "Crossbow Bolt",
	"Item_Ammo_Flare_C":                                           "Flare Gun Ammo",
	"Item_Ammo_Mortar_C":                                          "Mortar Ammo",
	"Item_Armor_C_01_Lv3_C":                                       "Military Vest (Level 3)",
	"Item_Armor_D_01_Lv2_C":                                       "Police Vest (Level 2)",
	"Item_Armor_E_01_Lv1_C":                                       "Police Vest (Level 1)",
	"Item_Attach_Weapon_Lower_AngledForeGrip_C":                   "Angled Foregrip",
	"Item_Attach_Weapon_Lower_Foregrip_C":                         "Vertical Foregrip",
	"Item_Attach_Weapon_Lower_HalfGrip_C":                         "Half Grip",
	"Item_Attach_Weapon_Lower_LaserPointer_C":                     "Laser Sight",
	"Item_Attach_Weapon_Lower_LightweightForeGrip_C":              "Light Grip",
	"Item_Attach_Weapon_Lower_QuickDraw_Large_Crossbow_C":         "QuickDraw Crossbow Quiver",
	"Item_Attach_Weapon_Lower_ThumbGrip_C":                        "Thumb Grip",
	"Item_Attach_Weapon_Magazine_ExtendedQuickDraw_Large_C":       "Extended QuickDraw Mag (AR, DMR, M249, S12K)",
	"Item_Attach_Weapon_Magazine_ExtendedQuickDraw_Medium_C":      "Extended QuickDraw Mag (Handgun, SMG)",
	"Item_Attach_Weapon_Magazine_ExtendedQuickDraw_Small_C":       "Extended QuickDraw Mag (Handgun)",
	"Item_Attach_Weapon_Magazine_ExtendedQuickDraw_SniperRifle_C": "Extended QuickDraw Mag (DMR, SR)",
	"Item_Attach_Weapon_Magazine_Extended_Large_C":                "Extended Mag (AR, DMR, M249, S12K)",
	"Item_Attach_Weapon_Magazine_Extended_Medium_C":               "Extended Mag (Handgun, SMG)",
	"Item_Attach_Weapon_Magazine_Extended_Small_C":                "Extended Mag (Handgun)",
	"Item_Attach_Weapon_Magazine_Extended_SniperRifle_C":          "Extended Mag (DMR, SR)",
	"Item_Attach_Weapon_Magazine_QuickDraw_Large_C":               "QuickDraw Mag (AR, DMR, M249, S12K)",
	"Item_Attach_Weapon_Magazine_QuickDraw_Medium_C":              "Quickdraw Mag (Handgun, SMG)",
	"Item_Attach_Weapon_Magazine_QuickDraw_Small_C":               "Quickdraw Mag (Handgun)",
	"Item_Attach_Weapon_Magazine_QuickDraw_SniperRifle_C":         "Quickdraw Mag (DMR, SR)",
	"Item_Attach_Weapon_Muzzle_Choke_C":                           "Choke",
	"Item_Attach_Weapon_Muzzle_Compensator_Large_C":               "Compensator (AR, DMR, S12K)",
	"Item_Attach_Weapon_Muzzle_Compensator_Medium_C":              "Compensator (Handgun, SMG)",
	"Item_Attach_Weapon_Muzzle_Compensator_SniperRifle_C":         "Compensator (DMR, SR)",
	"Item_Attach_Weapon_Muzzle_Duckbill_C":                        "Duckbill",
	"Item_Attach_Weapon_Muzzle_FlashHider_Large_C":                "Flash Hider (AR, DMR, S12K)",
	"Item_Attach_Weapon_Muzzle_FlashHider_Medium_C":               "Flash Hider (Handgun, SMG)",
	"Item_Attach_Weapon_Muzzle_FlashHider_SniperRifle_C":          "Flash Hider (DMR, SR)",
	"Item_Attach_Weapon_Muzzle_Suppressor_Large_C":                "Supressor (AR, DMR, S12K)",
	"Item_Attach_Weapon_Muzzle_Suppressor_Medium_C":               "Supressor (Handgun, SMG)",
	"Item_Attach_Weapon_Muzzle_Suppressor_Small_C":                "Supressor (Handgun)",
	"Item_Attach_Weapon_Muzzle_Suppressor_SniperRifle_C":          "Supressor (DMR, SR)",
	"Item_Attach_Weapon_SideRail_DotSight_RMR_C":                  "Canted Sight",
	"Item_Attach_Weapon_Stock_AR_Composite_C":                     "Tactical Stock",
	"Item_Attach_Weapon_Stock_AR_HeavyStock_C":                    "Heavy Stock",
	"Item_Attach_Weapon_Stock_Shotgun_BulletLoops_C":              "Shotgun Bullet Loops",
	"Item_Attach_Weapon_Stock_SniperRifle_BulletLoops_C":          "Sniper Rifle Bullet Loops",
	"Item_Attach_Weapon_Stock_SniperRifle_CheekPad_C":             "Sniper Rifle Cheek Pad",
	"Item_Attach_Weapon_Stock_UZI_C":                              "Uzi Stock",
	"Item_Attach_Weapon_Upper_ACOG_01_C":                          "4x ACOG Scope",
	"Item_Attach_Weapon_Upper_Aimpoint_C":                         "2x Aimpoint Scope",
	"Item_Attach_Weapon_Upper_CQBSS_C":                            "8x CQBSS Scope",
	"Item_Attach_Weapon_Upper_DotSight_01_C":                      "Red Dot Sight",
	"Item_Attach_Weapon_Upper_Holosight_C":                        "Holographic Sight",
	"Item_Attach_Weapon_Upper_PM2_01_C":                           "15x PM II Scope",
	"Item_Attach_Weapon_Upper_Scope3x_C":                          "3x Scope",
	"Item_Attach_Weapon_Upper_Scope6x_C":                          "6x Scope",
	"Item_Attach_Weapon_Upper_Thermal_C":                          "Thermal Scope",
	"Item_Back_B_01_StartParachutePack_C":                         "Parachute",
	"Item_Back_B_08_Lv3_C":                                        "Backpack (Level 3)",
	"Item_Back_BackupParachute_C":                                 "Emergency Parachute",
	"Item_Back_BlueBlocker":                                       "Jammer Pack",
	"Item_Back_C_01_Lv3_C":                                        "Backpack (Level 3)",
	"Item_Back_C_02_Lv3_C":                                        "Backpack (Level 3)",
	"Item_Back_E_01_Lv1_C":                                        "Backpack (Level 1)",
	"Item_Back_E_02_Lv1_C":                                        "Backpack (Level 1)",
	"Item_Back_F_01_Lv2_C":                                        "Backpack (Level 2)",
	"Item_Back_F_02_Lv2_C":                                        "Backpack (Level 2)",
	"Item_Boost_AdrenalineSyringe_C":                              "Adrenaline Syringe",
	"Item_Boost_EnergyDrink_C":                                    "Energy Drink",
	"Item_Boost_PainKiller_C":                                     "Painkiller",
	"Item_BulletproofShield_C":                                    "Folded Shield",
	"Item_Chimera_Key_C":                                          "Secret Room Key",
	"Item_DihorOtok_Key_C":                                        "Secret Room Key",
	"Item_EmergencyPickup_C":                                      "Emergency Pickup",
	"Item_Ghillie_01_C":                                           "Ghillie Suit",
	"Item_Ghillie_02_C":                                           "Ghillie Suit",
	"Item_Ghillie_03_C":                                           "Ghillie Suit",
	"Item_Ghillie_04_C":                                           "Ghillie Suit",
	"Item_Ghillie_05_C":                                           "Ghillie Suit",
	"Item_Ghillie_06_C":                                           "Ghillie Suit",
	"Item_Ghillie_07_C":                                           "Ghillie Suit",
	"Item_Head_E_01_Lv1_C":                                        "Motorcycle Helmet (Level 1)",
	"Item_Head_E_02_Lv1_C":                                        "Motorcycle Helmet (Level 1)",
	"Item_Head_F_01_Lv2_C":                                        "Military Helmet (Level 2)",
	"Item_Head_F_02_Lv2_C":                                        "Military Helmet (Level 2)",
	"Item_Head_G_01_Lv3_C":                                        "Spetsnaz Helmet (Level 3)",
	"Item_Heal_Bandage_C":                                         "Bandage",
	"Item_Heal_FirstAid_C":                                        "First Aid Kit",
	"Item_Heal_MedKit_C":                                          "Med kit",
	"Item_Heaven_Key_C":                                           "Key",
	"Item_JerryCan_C":                                             "Gas Can",
	"Item_Mountainbike_C":                                         "Mountain Bike",
	"Item_Secuity_KeyCard_C":                                      "Key Card",
	"Item_Special_Ascender_C":                                     "Ascender",
	"Item_Special_Ascender_NoChicken_C":                           "Ascender",
	"Item_Special_BackupParachute_C":                              "Backup Parachute",
	"Item_Tiger_Key_C":                                            "Secret Room Key",
	"Item_Tiger_SelfRevive_C":                                     "Self AED",
	"Item_Weapon_ACE32_C":                                         "ACE32",
	"Item_Weapon_AK47_C":                                          "AKM",
	"Item_Weapon_AUG_C":                                           "AUG A3",
	"Item_Weapon_AWM_C":                                           "AWM",
	"Item_Weapon_Apple_C":                                         "Apple",
	"Item_Weapon_Berreta686_C":                                    "S686",
	"Item_Weapon_BerylM762_C":                                     "Beryl",
	"Item_Weapon_BizonPP19_C":                                     "Bizon",
	"Item_Weapon_BlueChipDetector_C":                              "Blue Chip Detector",
	"Item_Weapon_BluezoneGrenade_C":                               "Bluezone Grenade",
	"Item_Weapon_C4_C":                                            "C4",
	"Item_Weapon_Cowbar_C":                                        "Crowbar",
	"Item_Weapon_Crossbow_C":                                      "Crossbow",
	"Item_Weapon_DP12_C":                                          "DBS",
	"Item_Weapon_DP28_C":                                          "DP-28",
	"Item_Weapon_DecoyGrenade_C":                                  "Decoy Grenade",
	"Item_Weapon_DesertEagle_C":                                   "Deagle",
	"Item_Weapon_Drone_C":                                         "Drone",
	"Item_Weapon_Duncans_M416_C":                                  "M416",
	"Item_Weapon_FNFal_C":                                         "SLR",
	"Item_Weapon_FlareGun_C":                                      "Flare Gun",
	"Item_Weapon_FlashBang_C":                                     "Flashbang",
	"Item_Weapon_G18_C":                                           "P18C",
	"Item_Weapon_G36C_C":                                          "G36C",
	"Item_Weapon_Grenade_C":                                       "Frag Grenade",
	"Item_Weapon_Grenade_Warmode_C":                               "Frag Grenade",
	"Item_Weapon_Groza_C":                                         "Groza",
	"Item_Weapon_HK416_C":                                         "M416",
	"Item_Weapon_K2_C":                                            "K2",
	"Item_Weapon_Kar98k_C":                                        "Kar98k",
	"Item_Weapon_L6_C":                                            "Lynx AMR",
	"Item_Weapon_M16A4_C":                                         "M16A4",
	"Item_Weapon_M1911_C":                                         "P1911",
	"Item_Weapon_M249_C":                                          "M249",
	"Item_Weapon_M24_C":                                           "M24",
	"Item_Weapon_M79_C":                                           "M79",
	"Item_Weapon_M9_C":                                            "P92",
	"Item_Weapon_MG3_C":                                           "MG3",
	"Item_Weapon_MP5K_C":                                          "MP5K",
	"Item_Weapon_MP9_C":                                           "MP9",
	"Item_Weapon_Machete_C":                                       "Machete",
	"Item_Weapon_Mads_QBU88_C":                                    "QBU88",
	"Item_Weapon_Mini14_C":                                        "Mini 14",
	"Item_Weapon_Mk12_C":                                          "Mk12",
	"Item_Weapon_Mk14_C":                                          "Mk14 EBR",
	"Item_Weapon_Mk47Mutant_C":                                    "Mk47 Mutant",
	"Item_Weapon_Molotov_C":                                       "Molotov Cocktail",
	"Item_Weapon_Mortar_C":                                        "Mortar",
	"Item_Weapon_Mosin_C":                                         "Mosin-Nagant",
	"Item_Weapon_NagantM1895_C":                                   "R1895",
	"Item_Weapon_OriginS12_C":                                     "O12",
	"Item_Weapon_P90_C":                                           "P90",
	"Item_Weapon_Pan_C":                                           "Pan",
	"Item_Weapon_PanzerFaust100M_C":                               "Panzerfaust",
	"Item_Weapon_QBU88_C":                                         "QBU88",
	"Item_Weapon_QBZ95_C":                                         "QBZ95",
	"Item_Weapon_Rhino_C":                                         "R45",
	"Item_Weapon_Rock_C":                                          "Rock",
	"Item_Weapon_SCAR-L_C":                                        "SCAR-L",
	"Item_Weapon_SKS_C":                                           "SKS",
	"Item_Weapon_Saiga12_C":                                       "S12K",
	"Item_Weapon_Sawnoff_C":                                       "Sawed-off",
	"Item_Weapon_Sickle_C":                                        "Sickle",
	"Item_Weapon_SmokeBomb_C":                                     "Smoke Grenade",
	"Item_Weapon_Snowball_C":                                      "Snowball",
	"Item_Weapon_SpikeTrap_C":                                     "Spike Trap",
	"Item_Weapon_Spotter_Scope_C":                                 "Spotter Scope",
	"Item_Weapon_StickyGrenade_C":                                 "Sticky Bomb",
	"Item_Weapon_TacPack_C":                                       "Tactical Pack",
	"Item_Weapon_Thompson_C":                                      "Tommy Gun",
	"Item_Weapon_TraumaBag_C":                                     "Trauma Bag",
	"Item_Weapon_UMP_C":                                           "UMP9",
	"Item_Weapon_UZI_C":                                           "Micro Uzi",
	"Item_Weapon_VSS_C":                                           "VSS",
	"Item_Weapon_Vector_C":                                        "Vector",
	"Item_Weapon_Win1894_C":                                       "Win94",
	"Item_Weapon_Winchester_C":                                    "S1897",
	"Item_Weapon_vz61Skorpion_C":                                  "Skorpion",
	"SP6_EventItem_DVD_01_C":                                      "Event Item",
	"SP6_EventItem_DVD_02_C":                                      "Event Item",
	"SP6_EventItem_DVD_03_C":                                      "Event Item",
	"Vehicle_Repair_Kit_C":                                        "Vehicle Repair Kit",
	"Vest_Repair_Kit_C":                                           "Vest Repair Kit",
	"WarModeStartParachutePack_C":                                 "Parachute",
}
