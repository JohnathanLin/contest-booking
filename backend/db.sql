CREATE TABLE `bookings`  (
     `booking_id` int(11) NOT NULL AUTO_INCREMENT COMMENT '预约id',
     `contest_id` int(11) NOT NULL COMMENT '比赛id',
     `contest_name` varchar(1024) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '比赛名称',
     `booking_status` int(11) NOT NULL COMMENT '比赛状态',
     `contest_time` datetime(0) NOT NULL COMMENT '比赛开打时间',
     `team_a` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'A队名称',
     `team_b` varchar(255) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT 'B队名称',
     `booking_score_a` int(11) NOT NULL COMMENT '预约的A队比分',
     `booking_score_b` int(11) NOT NULL COMMENT '预约的B队比分',
     `result_score_a` int(11) NOT NULL COMMENT 'A队结果比分',
     `result_score_b` int(11) NOT NULL COMMENT 'B队结果比分',
     `booking_email` varchar(4000) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL COMMENT '预约的邮箱信息',
     PRIMARY KEY (`booking_id`) USING BTREE
) ENGINE = InnoDB AUTO_INCREMENT = 8 CHARACTER SET = utf8mb4 COLLATE = utf8mb4_general_ci ROW_FORMAT = Dynamic;
