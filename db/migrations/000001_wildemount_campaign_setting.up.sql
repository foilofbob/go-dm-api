insert into campaign_setting (`name`)
values ("Wildemount");

select @wildemount_id:=id from campaign_setting where `name` = "Wildemount";

insert into `month` (campaign_setting_id, `name`, num_days, `order`)
values (@wildemount_id, "Horisal", 29, 1),
       (@wildemount_id, "Misuthar", 30, 2),
       (@wildemount_id, "Dualahei", 30, 3),
       (@wildemount_id, "Thunsheer", 31, 4),
       (@wildemount_id, "Unndilar", 28, 5),
       (@wildemount_id, "Brussendar", 31, 6),
       (@wildemount_id, "Sydenstar", 32, 7),
       (@wildemount_id, "Fessuran", 29, 8),
       (@wildemount_id, "Quen'pillar", 27, 9),
       (@wildemount_id, "Cuersaar", 29, 10),
       (@wildemount_id, "Duscar", 32, 11);

insert into week_day (campaign_setting_id, `name`, `order`)
values (@wildemount_id, "Miresen", 1),
       (@wildemount_id, "Grissen", 2),
       (@wildemount_id, "Whelsen", 3),
       (@wildemount_id, "Conthsen", 4),
       (@wildemount_id, "Folsen", 5),
       (@wildemount_id, "Yulisen", 6),
       (@wildemount_id, "Da'leysen", 7);

insert into calendar_cycle (campaign_setting_id, `name`, period)
values (@wildemount_id, "Catha", 35),
       (@wildemount_id, "Ruidus", 163);

insert into calendar_event (campaign_setting_id, month_id, day, name, description)
values (@wildemount_id, 1, 1, "New Dawn",
        "The first day of the new year is also the holy day of Avandra, the Changebringer, as the old year gives way to a new path. Emon celebrates New Dawn with a grand midnight feast, which commonly features a short play celebrating the changes witnessed in the past year. On the Menagerie Coast, New Dawn is celebrated with feasts on the shore at dusk the day before to watch the sunset. The feast continues through the night as celebrants discuss their hopes for the new year until the sun rises."),
       (@wildemount_id, 1, 27, "Hillsgold", "The celebration practices of Hillsgold are unknown."),
       (@wildemount_id, 2, 7, "Day of Challenging",
        "The holy day of Kord, the Stormlord is typically celebrated with a tournament called the Godsbrawl. The Day of Challenging is one of the most raucous holidays in Emon, and is also widely celebrated in Port Damali. Thousands of spectators attend the annual Godsbrawl, which is held in the fighting ring within each city's respective Temple of the Stormlord. The people root for their deity's favored champion, and there is a fierce (yet friendly) rivalry between the Champion of the Stormlord and the Champion of the Platinum Dragon. The winner earns the title of 'Supreme Champion' for an entire year. In Whitestone, the quadrennial Luncheon of Champions (which includes an all-out brawl) is also held on this day."),
       (@wildemount_id, 3, 13, "Renewal Festival", "The first day of spring is celebrated with fresh food, music, and games."),
       (@wildemount_id, 3, 20, "Wild's Grandeur",
        "Though the Arch Heart is the god of spring, the peak of the spring season is the holy day of Melora, the Wildmother. The people in the southern wilds of Tal'Dorei celebrate the Wildmother's strength by journeying to a place of great natural beauty. This could be the top of a mountainous waterfall, the center of a desert. Though Emon rarely celebrates Wild's Grandeur, the few who do will plant trees in observance of the holiday. The people of the Menagerie Coast take this day to sail for pleasure and observe the beauty of their surroundings. Those who partake in elements of Ki'Nau culture take the day to appreciate the fruits and foods given by the sea, and leave offering of delicacies and handmade crafts at altars of twisted roots and grasses. And in Vasselheim's Abundant Terrace, fresh vegetables are made into a variety of dishes, including succotash served from cauldrons."),
       (@wildemount_id, 4, 11, "Harvest's Rise",
        "The celebration practices of Harvest's Rise are unknown, but is presumably the beginning of the planting season."),
       (@wildemount_id, 4, 31, "Merryfrond's Day",
        "The celebration practices of Merryfrond's Day are unknown. Unofficially, this is like a Fey centered April Fools."),
       (@wildemount_id, 5, 18, "Deep Solace",
        "The holy day of Moradin, the All-Hammer is celebrated by especially devout followers in isolation. They meditate on the meaning of family and how they may be better mothers, fathers, siblings, and children. Dwarven communities, such as Kraghammer and Grimgolir, celebrate with a full day of feasting and drinking."),
       (@wildemount_id, 5, 26, "Zenith",
        "Summer begins on this day at high noon. It is celebrated with games, displays of magic, and black powder fireworks."),
       (@wildemount_id, 6, 15, "Artisan's Faire",
        "The celebration practices of Artisan's Faire are unknown, but is presumably a crafts festival."),
       (@wildemount_id, 6, 20, "Elvendawn, or Midsummer",
        "The holy day of the Corellon the Arch Heart celebrates the first emergence of the Elves into Exandria from the Feywild. In Syngorn, the elves open small doorways into the Feywild and celebrate alongside the wild fey with uncharacteristic vigor. The elves in Bysaes Tyl celebrate similarly to those in Syngorn, but it is a much quieter affair, as worship of Corellon is banned within the Dwendalian Empire. Still, the elves on this day often have a little more wine than usual."),
       (@wildemount_id, 7, 14, "Morn of Largesse",
        "The celebration practices of Morn of Largesse are unknown, but is presumably dedicated to charity."),
       (@wildemount_id, 7, 15, "Highsummer",
        "The holy day of Pelor, the Dawnfather is the peak of the summer season. Emon celebrates with an entire week of gift-giving and feasting, ending at midnight on the 21st of Sydenstar (the anniversary of the Battle of the Umbra Hills, where Zan Tal'Dorei dethroned Trist Drassig). Whitestone, where the Dawnfather is the city's patron god, celebrates with gift-giving and a festival of lights around the Sun Tree. When money was thin following the Briarwood occupation, most Whitestonians choose to recount the small things they are thankful for, rather than buy gifts. In Vasselheim's Abundant Terrace, which is ordinarily devoted to Melora, the celebration begins at dawn with a polenta made from freshly harvested corn and honey, spiced with chiles in grateful anticipation of the coming warmth of the sun. The Dwendalian Empire uses this day as an opportunity to promote enlistment to the Righteous Brand, which holds great feasts and hands out toy soldiers and other propaganda."),
       (@wildemount_id, 8, 3, "Harvest's Close",
        "Autumn begins on this day. Rural communities typically celebrate with feasts, and cities typically with carnivals."),
       (@wildemount_id, 9, 10, "The Hazel Festival",
        "The typical celebration practices of Hazel Festival are unknown, but the non-canonical 'Trinket's Honey Heist' (Sx27) takes place during the Hazel Festival in Westruun, where it has the atmosphere of a fair or autumn festival, with fried food, alcohol, and prizes for agricultural goods."),
       (@wildemount_id, 9, 22, "Civilization's Dawn",
        "The holy day of Erathis, the Lawbearer is the peak of the autumn season. Emon celebrates with a great bonfire in the square of each neighborhood, around which each community dances and gives gifts. In the Dwendalian Empire, people celebrate with feasts in honor of the Dwendal bloodline. One seat at each table is left open for the king, who eats in spirit with his subjects. In Vasselheim's Quadroads district, the Slayer's Take brings in great cuts of meat from their dangerous hunting trips into the Vesper Timberland and cooks them up in waist-high cooking pots for the community."),
       (@wildemount_id, 10, 13, "Night of Ascension",
        "Though the actual date of her rise to divinity is unclear, the holy day of The Raven Queen is nonetheless celebrated as the day of her apotheosis. Though most in Emon see this celebration of the dead as unnerving and macabre, the followers of the Matron of Ravens believe that the honored dead would prefer to be venerated with cheer, not misery. In the Dwendalian Empire, the Night of Ascension was once a similar cheery celebration of the dead, but it has become an occasion to burn effigies and decry the Kryn Dynasty for their unnatural relationship with death. In the Shadow Realm, beyond the Material Plane, the Night of Ascension is a major celebration showing the dark plane at its best."),
       (@wildemount_id, 10, 21, "Zan's Cup",
        "This holiday is only celebrated on Tal'Dorei. The celebration practices of Zan's Cup are unknown, but is presumably related to its namesake Zan Tal'Dorei."),
       (@wildemount_id, 11, 2, "Barren Eve",
        "Winter begins on this day, the longest night of the year. In Wildemount, it is a day of mourning for causalities in war, and many light candles to commemorate those lost."),
       (@wildemount_id, 11, 5, "Embertide",
        "The holy day of Bahamut, the Platinum Dragon is a day of remembrance. Solemnity and respect are shown to those who have fallen in the defense of others."),
       (@wildemount_id, 11, 16, "Day of Heart and Hearth",
        "A secular holiday celebrated in parts of Wildemount. In the Dwendalian Empire, it involves an outdoor celebration in the snow followed by an indoor celebration with family. On the Menagerie Coast, it's simply an excuse to party. It can be considered the Wildemount equivalent of Winter's Crest."),
       (@wildemount_id, 11, 20, "Winter's Crest",
        "This day celebrates the freedom of Tal'Dorei from Errevon the Rimelord at the end of the Icelost Years. It is often celebrated with omnipresent music in public areas, lavish gift-giving to relatives and loved ones, and the cutting and decorating of trees placed indoors. The Sun Tree in Whitestone is often decorated with lights and other baubles for Winter's Crest. Because it is the peak of the winter season, devout followers of the Matron of Ravens consider it to be one of her holy days, as she is the goddess of winter. Winter's Crest is also when the barrier between planes is at its thinnest. Because of this, some perform rituals concerning other planes on this day, such as the Briarwoods planning to perform their ritual within the Ziggurat beneath Whitestone to aid their master Vecna, or when Raishan was able to tear open Pyrah's rift to the Elemental Plane of Fire and allow Thordak back into Exandria.");
