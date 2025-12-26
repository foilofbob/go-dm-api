package domain

type Spell struct {
	ID             int
	Name           string
	Source         string
	Page           string
	Level          string
	CastingTime    string
	Duration       string
	School         string
	Range          string
	Components     string
	Classes        string
	VariantClasses string
	Subclasses     string
	Description    string
	HigherCasting  string
}

func ListSpells() ([]Spell, error) {
	rows, err := GetAll("spell")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var spells []Spell
	for rows.Next() {
		var spell Spell

		if err := rows.Scan(&spell.ID, &spell.Name, &spell.Source, &spell.Page, &spell.Level, &spell.CastingTime, &spell.Duration, &spell.School, &spell.Range, &spell.Components, &spell.Classes, &spell.VariantClasses, &spell.Subclasses, &spell.Description, &spell.HigherCasting); err != nil {
			return spells, err
		}

		spells = append(spells, spell)
	}

	if err = rows.Err(); err != nil {
		return spells, err
	}

	return spells, nil
}
