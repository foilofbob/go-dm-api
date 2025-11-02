package domain

import (
	"github.com/shopspring/decimal"
)

type Character struct {
	ID                int
	CampaignID        int
	Name              string
	Race              string
	Class             string
	ArmorClass        int
	HitPoints         int
	PassivePerception int
	Languages         string
	Movement          int
	Strength          int
	Dexterity         int
	Constitution      int
	Intelligence      int
	Wisdom            int
	Charisma          int
	Proficiencies     string
	PlayerType        string

	StrengthSaveProficiency     bool
	DexteritySaveProficiency    bool
	ConstitutionSaveProficiency bool
	IntelligenceSaveProficiency bool
	WisdomSaveProficiency       bool
	CharismaSaveProficiency     bool

	AcrobaticsProficiencyBonus     decimal.Decimal
	AnimalHandlingProficiencyBonus decimal.Decimal
	ArcanaProficiencyBonus         decimal.Decimal
	AthleticsProficiencyBonus      decimal.Decimal
	DeceptionProficiencyBonus      decimal.Decimal
	HistoryProficiencyBonus        decimal.Decimal
	InsightProficiencyBonus        decimal.Decimal
	IntimidationProficiencyBonus   decimal.Decimal
	InvestigationProficiencyBonus  decimal.Decimal
	MedicineProficiencyBonus       decimal.Decimal
	NatureProficiencyBonus         decimal.Decimal
	PerceptionProficiencyBonus     decimal.Decimal
	PerformanceProficiencyBonus    decimal.Decimal
	PersuasionProficiencyBonus     decimal.Decimal
	ReligionProficiencyBonus       decimal.Decimal
	SleightOfHandProficiencyBonus  decimal.Decimal
	StealthProficiencyBonus        decimal.Decimal
	SurvivalProficiencyBonus       decimal.Decimal
}

func GetCharacter(id int) (*Character, error) {
	row := GetByID("character", id)
	character := &Character{}
	readErr := row.Scan(&character.ID, &character.CampaignID, &character.Name, &character.Race, &character.Class,
		&character.ArmorClass, &character.HitPoints, &character.PassivePerception, &character.Languages,
		&character.Movement, &character.Strength, &character.Dexterity, &character.Constitution, &character.Intelligence,
		&character.Wisdom, &character.Charisma, &character.Proficiencies, &character.StrengthSaveProficiency,
		&character.DexteritySaveProficiency, &character.ConstitutionSaveProficiency, &character.IntelligenceSaveProficiency,
		&character.WisdomSaveProficiency, &character.CharismaSaveProficiency, &character.AcrobaticsProficiencyBonus,
		&character.AnimalHandlingProficiencyBonus, &character.ArcanaProficiencyBonus, &character.AthleticsProficiencyBonus,
		&character.DeceptionProficiencyBonus, &character.HistoryProficiencyBonus, &character.InsightProficiencyBonus,
		&character.IntimidationProficiencyBonus, &character.InvestigationProficiencyBonus, &character.MedicineProficiencyBonus,
		&character.NatureProficiencyBonus, &character.PerceptionProficiencyBonus, &character.PerformanceProficiencyBonus,
		&character.PersuasionProficiencyBonus, &character.ReligionProficiencyBonus, &character.SleightOfHandProficiencyBonus,
		&character.StealthProficiencyBonus, &character.SurvivalProficiencyBonus, &character.PlayerType)

	if readErr != nil {
		return nil, readErr
	}
	return character, nil
}

func GetCharacters(campaignId int, playerType string) ([]Character, error) {
	query := "SELECT * FROM `character` WHERE campaign_id = ? AND player_type = ?"
	rows, err := DBQuery(query, campaignId, playerType)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var characters []Character

	for rows.Next() {
		var character Character

		if err := rows.Scan(&character.ID, &character.CampaignID, &character.Name, &character.Race, &character.Class,
			&character.ArmorClass, &character.HitPoints, &character.PassivePerception, &character.Languages,
			&character.Movement, &character.Strength, &character.Dexterity, &character.Constitution, &character.Intelligence,
			&character.Wisdom, &character.Charisma, &character.Proficiencies, &character.StrengthSaveProficiency,
			&character.DexteritySaveProficiency, &character.ConstitutionSaveProficiency, &character.IntelligenceSaveProficiency,
			&character.WisdomSaveProficiency, &character.CharismaSaveProficiency, &character.AcrobaticsProficiencyBonus,
			&character.AnimalHandlingProficiencyBonus, &character.ArcanaProficiencyBonus, &character.AthleticsProficiencyBonus,
			&character.DeceptionProficiencyBonus, &character.HistoryProficiencyBonus, &character.InsightProficiencyBonus,
			&character.IntimidationProficiencyBonus, &character.InvestigationProficiencyBonus, &character.MedicineProficiencyBonus,
			&character.NatureProficiencyBonus, &character.PerceptionProficiencyBonus, &character.PerformanceProficiencyBonus,
			&character.PersuasionProficiencyBonus, &character.ReligionProficiencyBonus, &character.SleightOfHandProficiencyBonus,
			&character.StealthProficiencyBonus, &character.SurvivalProficiencyBonus, &character.PlayerType); err != nil {
			return characters, err
		}

		characters = append(characters, character)
	}

	if err = rows.Err(); err != nil {
		return characters, err
	}

	return characters, nil
}

func GetPlayerCharacters(campaignId int) ([]Character, error) {
	return GetCharacters(campaignId, "PLAYER")
}

func GetNonPlayerCharacters(campaignId int) ([]Character, error) {
	return GetCharacters(campaignId, "NPC")
}

func CreateCharacter(campaignID int, name string, race string, class string, armorClass int, hitPoints int,
	passivePerception int, languages string, movement int, strength int, dexterity int, constitution int,
	intelligence int, wisdom int, charisma int, proficiencies string, playerType string, strengthSaveProficiency bool,
	dexteritySaveProficiency bool, constitutionSaveProficiency bool, intelligenceSaveProficiency bool,
	wisdomSaveProficiency bool, charismaSaveProficiency bool, acrobaticsProficiencyBonus decimal.Decimal,
	animalHandlingProficiencyBonus decimal.Decimal, arcanaProficiencyBonus decimal.Decimal, athleticsProficiencyBonus decimal.Decimal,
	deceptionProficiencyBonus decimal.Decimal, historyProficiencyBonus decimal.Decimal, insightProficiencyBonus decimal.Decimal,
	intimidationProficiencyBonus decimal.Decimal, investigationProficiencyBonus decimal.Decimal, medicineProficiencyBonus decimal.Decimal,
	natureProficiencyBonus decimal.Decimal, perceptionProficiencyBonus decimal.Decimal, performanceProficiencyBonus decimal.Decimal,
	persuasionProficiencyBonus decimal.Decimal, religionProficiencyBonus decimal.Decimal, sleightOfHandProficiencyBonus decimal.Decimal,
	stealthProficiencyBonus decimal.Decimal, survivalProficiencyBonus decimal.Decimal) (*Character, error) {
	db := DBConnection()
	defer db.Close()

	query := `
		INSERT INTO ` + "`character`" + `
		    (campaign_id, name, race, class, armor_class, hit_points, passive_perception, languages,
		                       movement, strength, dexterity, constitution, intelligence, wisdom, charisma,
		                       proficiencies, player_type, strength_save_proficiency, dexterity_save_proficiency, 
		                       constitution_save_proficiency, intelligence_save_proficiency, wisdom_save_proficiency, 
		                       charisma_save_proficiency, acrobatics_proficiency_bonus, animal_handling_proficiency_bonus, 
		                       arcana_proficiency_bonus, athletics_proficiency_bonus, deception_proficiency_bonus, 
		                       history_proficiency_bonus, insight_proficiency_bonus, intimidation_proficiency_bonus, 
		                       investigation_proficiency_bonus, medicine_proficiency_bonus, nature_proficiency_bonus, 
		                       perception_proficiency_bonus, performance_proficiency_bonus, persuasion_proficiency_bonus, 
		                       religion_proficiency_bonus, sleight_of_hand_proficiency_bonus, stealth_proficiency_bonus, 
		                       survival_proficiency_bonus)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
	`

	res, insertErr := db.Exec(query, campaignID, name, race, class, armorClass, hitPoints, passivePerception, languages,
		movement, strength, dexterity, constitution, intelligence, wisdom, charisma, proficiencies, playerType,
		strengthSaveProficiency, dexteritySaveProficiency, constitutionSaveProficiency, intelligenceSaveProficiency,
		wisdomSaveProficiency, charismaSaveProficiency, acrobaticsProficiencyBonus, animalHandlingProficiencyBonus,
		arcanaProficiencyBonus, athleticsProficiencyBonus, deceptionProficiencyBonus, historyProficiencyBonus,
		insightProficiencyBonus, intimidationProficiencyBonus, investigationProficiencyBonus, medicineProficiencyBonus,
		natureProficiencyBonus, perceptionProficiencyBonus, performanceProficiencyBonus, persuasionProficiencyBonus,
		religionProficiencyBonus, sleightOfHandProficiencyBonus, stealthProficiencyBonus, survivalProficiencyBonus)
	if insertErr != nil {
		return nil, insertErr
	}

	lid, err := res.LastInsertId()
	if err != nil {
		return nil, insertErr
	}

	return GetCharacter(int(lid))
}

func UpdateCharacter(characterID int, name string, race string, class string, armorClass int, hitPoints int,
	passivePerception int, languages string, movement int, strength int, dexterity int, constitution int,
	intelligence int, wisdom int, charisma int, proficiencies string, playerType string, strengthSaveProficiency bool,
	dexteritySaveProficiency bool, constitutionSaveProficiency bool, intelligenceSaveProficiency bool,
	wisdomSaveProficiency bool, charismaSaveProficiency bool, acrobaticsProficiencyBonus decimal.Decimal,
	animalHandlingProficiencyBonus decimal.Decimal, arcanaProficiencyBonus decimal.Decimal, athleticsProficiencyBonus decimal.Decimal,
	deceptionProficiencyBonus decimal.Decimal, historyProficiencyBonus decimal.Decimal, insightProficiencyBonus decimal.Decimal,
	intimidationProficiencyBonus decimal.Decimal, investigationProficiencyBonus decimal.Decimal, medicineProficiencyBonus decimal.Decimal,
	natureProficiencyBonus decimal.Decimal, perceptionProficiencyBonus decimal.Decimal, performanceProficiencyBonus decimal.Decimal,
	persuasionProficiencyBonus decimal.Decimal, religionProficiencyBonus decimal.Decimal, sleightOfHandProficiencyBonus decimal.Decimal,
	stealthProficiencyBonus decimal.Decimal, survivalProficiencyBonus decimal.Decimal) (*Character, error) {
	db := DBConnection()
	defer db.Close()

	query := `
		UPDATE ` + "`character`" + `
		SET name = ?, race = ?, class = ?, armor_class = ?, hit_points = ?, passive_perception = ?, languages = ?, 
		    movement = ?, strength = ? , dexterity = ? , constitution = ? , intelligence = ?, wisdom = ? , charisma = ?, 
		    proficiencies = ?, player_type = ?, strength_save_proficiency = ?, dexterity_save_proficiency = ?, 
		    constitution_save_proficiency = ?, intelligence_save_proficiency = ?, wisdom_save_proficiency = ?, 
		    charisma_save_proficiency = ?, acrobatics_proficiency_bonus = ?, 
		    animal_handling_proficiency_bonus = ?, arcana_proficiency_bonus = ?, athletics_proficiency_bonus = ?, 
		    deception_proficiency_bonus = ?, history_proficiency_bonus = ?, insight_proficiency_bonus = ?, 
		    intimidation_proficiency_bonus = ?, investigation_proficiency_bonus = ?, medicine_proficiency_bonus = ?, 
		    nature_proficiency_bonus = ?, perception_proficiency_bonus = ?, performance_proficiency_bonus = ?, 
		    persuasion_proficiency_bonus = ?, religion_proficiency_bonus = ?, sleight_of_hand_proficiency_bonus = ?, 
		    stealth_proficiency_bonus = ?, survival_proficiency_bonus = ? 
		WHERE id = ?
	`

	_, updateErr := db.Exec(query, name, race, class, armorClass, hitPoints, passivePerception, languages, movement,
		strength, dexterity, constitution, intelligence, wisdom, charisma, proficiencies, playerType, strengthSaveProficiency,
		dexteritySaveProficiency, constitutionSaveProficiency, intelligenceSaveProficiency, wisdomSaveProficiency,
		charismaSaveProficiency, acrobaticsProficiencyBonus, animalHandlingProficiencyBonus, arcanaProficiencyBonus,
		athleticsProficiencyBonus, deceptionProficiencyBonus, historyProficiencyBonus, insightProficiencyBonus,
		intimidationProficiencyBonus, investigationProficiencyBonus, medicineProficiencyBonus, natureProficiencyBonus,
		perceptionProficiencyBonus, performanceProficiencyBonus, persuasionProficiencyBonus, religionProficiencyBonus,
		sleightOfHandProficiencyBonus, stealthProficiencyBonus, survivalProficiencyBonus, characterID)
	if updateErr != nil {
		return nil, updateErr
	}
	return GetCharacter(characterID)
}

func DeleteCharacter(characterID int) error {
	db := DBConnection()
	defer db.Close()

	query := "DELETE FROM `character` WHERE id = ?"
	res, deleteErr := db.Exec(query, characterID)

	if deleteErr != nil {
		println("Delete character error: " + deleteErr.Error())
	}

	println(res.RowsAffected())

	return deleteErr
}
