truncate table point_of_interest;
drop table point_of_interest;
delete from note where reference_type = "POINT_OF_INTEREST";

truncate table sublocation;
drop table sublocation;

truncate table location;
drop table location;
delete from note where reference_type = "LOCATION";
