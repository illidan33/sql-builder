package sql_builder

import (
	"testing"
)

type Skill struct {
	Condition string `json:"condition" db:"condition"`
	Desc      string `json:"desc" db:"desc"`
	SkillType int    `json:"skillType" db:"skill_type"`
	Status    int    `json:"status" db:"status"`
	Empty     string `json:"empty"`
}

func TestSelectSqlBuilder_SelectByStruct(t *testing.T) {
	skill := Skill{
		Condition: "",
		Desc:      "test",
		SkillType: 1,
		Status:    0,
	}

	builder := Select("skill")
	builder.SelectByStruct(skill, true)

	if builder.String() != "SELECT `condition`,`desc`,`skill_type`,`status` FROM `skill` WHERE `desc`=? AND `skill_type`=?;" {
		t.Fatalf("Sql Error: %s \n", builder.String())
	}
	args := builder.Args()
	if args[0] != "test" || args[1] != 1 {
		t.Errorf("Args Error: %+v\n", args)
	}
}

func TestSelect(t *testing.T) {
	builder := Select("skill")
	builder.SetSearchFields("`condition`,`skill_type`")
	builder.WhereEq("skill_type", 1)
	builder.WhereGt("skill_type", 5)
	builder.WhereLt("skill_type", 10)
	builder.WhereIn("status", []interface{}{1, 2})
	builder.WhereLike("condition", "test condition")
	builder.WhereOr([]WhereOrCondition{
		{
			FieldName:  "desc",
			WhereType:  WHERE_TYPE_LikeAFT,
			FieldValue: "test desc",
		}, {
			FieldName:  "desc",
			WhereType:  WHERE_TYPE_LikeBEF,
			FieldValue: "test desc",
		},
	})
	// 注意一下使用顺序
	builder.GroupBy("`id`,`skill_type`")
	builder.OrderBy("`skill_type` ASC")
	builder.Limit(0, 20)

	if builder.String() != "SELECT `condition`,`skill_type` FROM `skill` WHERE (`skill_type`=? AND `skill_type`>? AND `skill_type`<? AND status IN (?,?) AND `condition` LIKE ?) OR (`desc` LIKE ? AND `desc` LIKE ?) GROUP BY `id`,`skill_type` ORDER BY `skill_type` ASC LIMIT 0,20;" {
		t.Errorf("Sql Error: %s \n", builder.String())
	}

	args := builder.Args()
	if args[0] != 1 || args[1] != 5 || args[2] != 10 || args[3] != 1 || args[4] != 2 || args[5] != "%test condition%" || args[6] != "test desc%" || args[7] != "%test desc" {
		t.Errorf("Args Error: %+v \n", args)
	}
}

func TestUpdateSqlBuilder_UpdateByStruct(t *testing.T) {
	skill := Skill{
		Condition: "",
		Desc:      "test",
		SkillType: 1,
		Status:    0,
	}

	builder := Update("skill")
	builder.UpdateByStruct(skill, true)
	builder.WhereEq("skill_type", 1)
	builder.WhereIn("status", []interface{}{1, 2})
	builder.WhereLike("condition", "test")

	if builder.String() != "UPDATE `skill` SET `desc`=?,`skill_type`=? WHERE `skill_type`=? AND status IN (?,?) AND `condition` LIKE ?;" {
		t.Errorf("Sql Error: %s \n", builder.String())
	}

	args := builder.Args()
	if args[0] != "test" || args[1] != 1 || args[2] != 1 || args[3] != 1 || args[4] != 2 || args[5] != "%test%" {
		t.Errorf("Args Error: %+v \n", args)
	}
}

func TestInsertSqlBuilder_InsertByStruct(t *testing.T) {
	skill := Skill{
		Condition: "",
		Desc:      "test",
		SkillType: 1,
	}

	builder := Insert("skill")
	builder.InsertByStruct(skill)

	if builder.String() != "INSERT INTO `skill`(`condition`,`desc`,`skill_type`,`status`) VALUES(?,?,?,?);" {
		t.Errorf("Sql Error: %s \n", builder.String())
	}

	args := builder.Args()
	if args[0] != "" || args[1] != "test" || args[2] != 1 || args[3] != 0 {
		t.Errorf("Args Error: %+v", args)
	}
}

func TestDeleteSqlBuilder_DeleteAndWhereByStruct(t *testing.T) {
	skill := Skill{
		SkillType: 1,
	}

	builder := Delete("skill")
	builder.WhereByStruct(skill, true)

	if builder.String() != "DELETE FROM `skill` WHERE `skill_type`=?;" {
		t.Fatalf("Error -- sql string: %s \n", builder.String())
	}
}
