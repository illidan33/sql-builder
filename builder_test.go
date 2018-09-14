package sql_builder

import (
	"testing"
)

type Skill struct {
	Condition string `json:"condition" db:"condition"`
	Desc      string `json:"desc" db:"desc"`
	SkillType int    `json:"skillType" db:"skill_type"`
}

func TestSelectSqlBuilder_SelectByStruct(t *testing.T) {
	skill := Skill{
		SkillType: 1,
	}

	builder := Select("skill")
	builder.SelectByStruct(skill)

	if builder.String() != "SELECT `condition`,`desc`,`skill_type` FROM `skill` WHERE `skill_type`=?;" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}

func TestUpdateSqlBuilder_UpdateByStruct(t *testing.T) {
	skill := Skill{
		SkillType: 1,
	}

	builder := Update("skill")
	builder.UpdateByStruct(skill, true)

	if builder.String() != "UPDATE `skill` SET `skill_type`=?;" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}

func TestInsertSqlBuilder_InsertByStruct(t *testing.T) {
	skill := Skill{
		SkillType: 1,
	}

	builder := Insert("skill")
	builder.InsertByStruct(skill)

	if builder.String() != "INSERT INTO `skill`(condition,desc,skill_type) VALUES(?,?,?);" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}

func TestDeleteSqlBuilder_DeleteAndWhereByStruct(t *testing.T) {
	skill := Skill{
		SkillType: 1,
	}

	builder := Delete("skill")
	builder.WhereByStruct(skill)

	if builder.String() != "DELETE FROM `skill` WHERE `skill_type`=?;" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}
