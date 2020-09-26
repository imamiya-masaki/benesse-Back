-- INSERT INTO collectq_db.quetions (title, content) VALUES 
-- ("gitむずかちい", "gitってなんであんなに難しいんですか?\n簡単にしてほしい"), ("docker is god", "dockerって神ですね^^\n簡単に初期化できるサイト知ってますか?"), ("好きなアニメを教えてください", "好きなアニメを教えてください。\n私は宇宙パトロールです");

INSERT INTO collectq_db.users (login_id, login_pass, benesse_id) VALUES
("test1","pass1","benesse1"),("test2","pass2","benesse2"),("test3","pass3","benesse3");

INSERT INTO collectq_db.user_stamps (user_id, stamp_count, stamp_type) VALUES
(1,3,1),(1,5,2),(1,10,3),(2,3,1),(2,5,2),(2,10,3),(3,3,1),(3,5,2),(3,10,3);


INSERT INTO collectq_db.user_stus (user_id, japanese_score, society_score, math_score, science_score, english_score) VALUES
(1,30,50,60,30,80),(2,80,10,0,80,100),(3,25,35,43,21,11);

INSERT INTO collectq_db.items (stamp_type, stamp_count,  item_name) VALUES
(1,15,"小説"),(2,15,"靴"),(3,15,"チョコレート"),(1,30,"百科事典"),(2,30,"ローラーシューズ"),(3,30,"ケーキ");