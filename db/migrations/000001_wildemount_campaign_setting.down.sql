select @wildemount_id:=id from campaign_setting where `name` = "Wildemount";

delete from calendar_event where campaign_setting_id = @wildemount_id;
delete from `month` where campaign_setting_id = @wildemount_id;
delete from week_day where campaign_setting_id = @wildemount_id;
delete from calendar_cycle where campaign_setting_id = @wildemount_id;
delete from campaign_setting where campaign_setting_id = @wildemount_id;
