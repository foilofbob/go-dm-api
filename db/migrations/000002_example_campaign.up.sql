insert into campaign (`name`, current_player_xp, campaign_setting_id)
values ("My First Campaign", 0, 1);

select @campaign_id := id
from campaign
where `name` = "My First Campaign";

select @wildemount_id := id
from campaign_setting
where `name` = "Wildemount";

select @catha_id := id
from calendar_cycle
where campaign_setting_id = @wildemount_id
  and `name` = "Catha";

select @ruidus_id := id
from calendar_cycle
where campaign_setting_id = @wildemount_id
  and `name` = "Ruidus";

insert into campaign_calendar_cycle_offset (campaign_id, calendar_cycle_id, offset)
values (@campaign_id, @catha_id, floor(rand() * 34)),
       (@campaign_id, @ruidus_id, floor(rand() * 164));

insert into game_day (campaign_id, in_game_day, `day`, `month`, `year`)
values (@campaign_id, 1, floor(rand() * 28 + 1), floor(rand() * 12 + 1), 837);

select @game_day_id := id
from game_day
where campaign_id = @campaign_id;

insert into note (campaign_id, reference_type, reference_id, title, content)
values (@campaign_id, "GAME_DAY", @game_day_id, "Day", "!!! info This is all markdown! Lots of stuff you can do here.");
