select @campaign_id := id
from campaign
where `name` = "My First Campaign";

delete
from note
where campaign_id = @campaign_id;

delete
from game_day
where campaign_id = @campaign_id;

delete
from campaign_calendar_cycle_offset
where campaign_id = @campaign_id;

delete
from campaign
where id = @campaign_id;
