# rename to reflect a structure that will apply for both PCs and NPCs
alter table player rename `character`;

# decided we will want to track specific saving throw proficiency
alter table `character`
    drop column saves,
    add column strength_save_proficiency boolean default false,
    add column dexterity_save_proficiency boolean default false,
    add column constitution_save_proficiency boolean default false,
    add column intelligence_save_proficiency boolean default false,
    add column wisdom_save_proficiency boolean default false,
    add column charisma_save_proficiency boolean default false;

# add columns to track specific proficiency bonuses - might need overrides, like for passive_perception
alter table `character`
    add column acrobatics_proficiency_bonus decimal(2,1) default 0.0,
    add column animal_handling_proficiency_bonus decimal(2,1) default 0.0,
    add column arcana_proficiency_bonus decimal(2,1) default 0.0,
    add column athletics_proficiency_bonus decimal(2,1) default 0.0,
    add column deception_proficiency_bonus decimal(2,1) default 0.0,
    add column history_proficiency_bonus decimal(2,1) default 0.0,
    add column insight_proficiency_bonus decimal(2,1) default 0.0,
    add column intimidation_proficiency_bonus decimal(2,1) default 0.0,
    add column investigation_proficiency_bonus decimal(2,1) default 0.0,
    add column medicine_proficiency_bonus decimal(2,1) default 0.0,
    add column nature_proficiency_bonus decimal(2,1) default 0.0,
    add column perception_proficiency_bonus decimal(2,1) default 0.0,
    add column performance_proficiency_bonus decimal(2,1) default 0.0,
    add column persuasion_proficiency_bonus decimal(2,1) default 0.0,
    add column religion_proficiency_bonus decimal(2,1) default 0.0,
    add column sleight_of_hand_proficiency_bonus decimal(2,1) default 0.0,
    add column stealth_proficiency_bonus decimal(2,1) default 0.0,
    add column survival_proficiency_bonus decimal(2,1) default 0.0;

alter table `character`
    add column player_type varchar(20) not null default 'PLAYER';
