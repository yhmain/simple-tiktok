/*
 Navicat Premium Data Transfer

 Source Server         : Simple-TikTok
 Source Server Type    : MySQL
 Source Server Version : 80028
 Source Host           : localhost:3306
 Source Schema         : tiktok

 Target Server Type    : MySQL
 Target Server Version : 80028
 File Encoding         : 65001

 Date: 17/05/2022 17:38:45
*/

SET NAMES utf8mb4;
SET FOREIGN_KEY_CHECKS = 0;

-- ----------------------------
-- Table structure for comments
-- ----------------------------
DROP TABLE IF EXISTS `comments`;
CREATE TABLE `comments`  (
  `Cid` int NOT NULL COMMENT '评论id',
  `Content` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '评论内容',
  `CreateDate` bigint NULL DEFAULT NULL COMMENT '评论创建时间',
  `Uid` bigint NULL DEFAULT NULL COMMENT '评论的用户id',
  `Vid` int NULL DEFAULT NULL COMMENT '评论的视频id',
  PRIMARY KEY (`Cid`) USING BTREE,
  INDEX `User Comment`(`Uid`) USING BTREE,
  INDEX `Video Comment`(`Vid`) USING BTREE,
  CONSTRAINT `User Comment` FOREIGN KEY (`Uid`) REFERENCES `users` (`Uid`) ON DELETE RESTRICT ON UPDATE RESTRICT,
  CONSTRAINT `Video Comment` FOREIGN KEY (`Vid`) REFERENCES `videos` (`Vid`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of comments
-- ----------------------------
INSERT INTO `comments` VALUES (1, '测试评论1', 1652773376, 2, 2);

-- ----------------------------
-- Table structure for users
-- ----------------------------
DROP TABLE IF EXISTS `users`;
CREATE TABLE `users`  (
  `Uid` bigint NOT NULL COMMENT '用户id，需要唯一（抖音号）',
  `NickName` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '用户名：登录时用，需要唯一',
  `UserPwd` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '秘密：登录时用',
  `FollowCount` int NULL DEFAULT 0 COMMENT '关注数',
  `FollowerCount` int NULL DEFAULT 0 COMMENT '粉丝数',
  `IsFollow` tinyint UNSIGNED NULL DEFAULT 0 COMMENT 'true:已关注，false:未关注',
  `CommentCount` int UNSIGNED NULL DEFAULT 0 COMMENT '评论数目',
  `IsFavorite` tinyint NULL DEFAULT 0 COMMENT '是否喜欢',
  `FavoriteCount` int NULL DEFAULT 0 COMMENT '喜欢的人数',
  PRIMARY KEY (`Uid`) USING BTREE
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of users
-- ----------------------------
INSERT INTO `users` VALUES (1, 'zhanglei', 'douyin', 10, 10, 0, 0, 0, 0);
INSERT INTO `users` VALUES (2, 'admin', '123456', 20, 20, 0, 0, 0, 0);
INSERT INTO `users` VALUES (3, 'yuhang', '123456', 30, 30, 0, 0, 0, 0);
INSERT INTO `users` VALUES (4, 'qinyuan', '123456', 40, 40, 0, 0, 0, 0);
INSERT INTO `users` VALUES (5, 'wangerke', '123456', 50, 50, 0, 0, 0, 0);
INSERT INTO `users` VALUES (6, 'zhangzhuoxun', '123456', 60, 60, 0, 0, 0, 0);
INSERT INTO `users` VALUES (7, 'guzhongqing', '123456', 70, 70, 0, 0, 0, 0);
INSERT INTO `users` VALUES (4816802046802006016, 'yh@qq.com', '77899900', 0, 0, 0, 0, 0, 0);
INSERT INTO `users` VALUES (5450728872746291200, 'eqq@163.com', '123456', 0, 0, 0, 0, 0, 0);

-- ----------------------------
-- Table structure for videos
-- ----------------------------
DROP TABLE IF EXISTS `videos`;
CREATE TABLE `videos`  (
  `Vid` int NOT NULL COMMENT '视频id，唯一',
  `VideoTitle` text CHARACTER SET utf8 COLLATE utf8_general_ci NULL COMMENT '视频内容：描述',
  `PlayUrl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频播放地址',
  `CoverUrl` varchar(255) CHARACTER SET utf8 COLLATE utf8_general_ci NULL DEFAULT NULL COMMENT '视频封面地址',
  `FavoriteCount` int NULL DEFAULT NULL COMMENT '视频点赞总数',
  `CommentCount` int NULL DEFAULT NULL COMMENT '视频评论总数',
  `IsFavorite` tinyint(1) NULL DEFAULT NULL COMMENT 'true:已点赞，false:未点赞',
  `CreatedTime` bigint NULL DEFAULT NULL COMMENT '视频发布时间，存储时间戳',
  `Uid` bigint NULL DEFAULT NULL COMMENT '外键：用户ID',
  PRIMARY KEY (`Vid`) USING BTREE,
  INDEX `User Video`(`Uid`) USING BTREE,
  CONSTRAINT `User Video` FOREIGN KEY (`Uid`) REFERENCES `users` (`Uid`) ON DELETE RESTRICT ON UPDATE RESTRICT
) ENGINE = InnoDB CHARACTER SET = utf8 COLLATE = utf8_general_ci ROW_FORMAT = Dynamic;

-- ----------------------------
-- Records of videos
-- ----------------------------
INSERT INTO `videos` VALUES (1, '这是第一条视频的内容介绍：熊熊真可爱！', 'bear.mp4', 'bear.jpg', 10, 10, 0, 1652590561, 1);
INSERT INTO `videos` VALUES (2, '这是第二条视频的内容介绍：奶龙作为封面！', 'maxclub.mp4', 'long.jpeg', 5, 5, 0, 1652676961, 1);
INSERT INTO `videos` VALUES (3, '这是第三条视频的内容介绍：熊熊作为封面！', 'bear.mp4', 'long.jpeg', 3, 3, 0, 1652675588, 3);
INSERT INTO `videos` VALUES (11, '视频内容：popo猫', 'popo猫.mp4', 'popo猫.jpg', 10, 10, 0, 1652590561, 1);
INSERT INTO `videos` VALUES (12, '视频内容：rickroll', 'rickroll.mp4', 'rickroll.jpg', 10, 10, 0, 1652590571, 2);
INSERT INTO `videos` VALUES (13, '视频内容：尬笑', '尬笑.mp4', '尬笑.jpg', 10, 10, 0, 1652590581, 3);
INSERT INTO `videos` VALUES (14, '视频内容：狗子跳舞', '狗子跳舞.mp4', '狗子跳舞.jpg', 10, 10, 0, 1652590591, 4);
INSERT INTO `videos` VALUES (15, '视频内容：黑人落泪', '黑人落泪.mp4', '黑人落泪.jpg', 10, 10, 0, 1652590601, 5);
INSERT INTO `videos` VALUES (16, '视频内容：肯德基老爷爷', '肯德基老爷爷.mp4', '肯德基老爷爷.jpg', 10, 10, 0, 1652590601, 6);
INSERT INTO `videos` VALUES (17, '视频内容：老板笑', '老板笑.mp4', '老板笑.jpg', 10, 10, 0, 1652590611, 7);
INSERT INTO `videos` VALUES (18, '视频内容：你怎么睡得着的', '你怎么睡得着的.mp4', '你怎么睡得着的.jpg', 10, 10, 0, 1652590621, 1);
INSERT INTO `videos` VALUES (19, '视频内容：这一脚踢出了整个盛夏', '这一脚踢出了整个盛夏.mp4', '这一脚踢出了整个盛夏.jpg', 10, 10, 0, 1652590631, 2);
INSERT INTO `videos` VALUES (20, '视频内容：上帝DJ打碟', '上帝DJ打碟.mp4', '上帝DJ打碟.jpg', 10, 10, 0, 1652590641, 3);
INSERT INTO `videos` VALUES (21, '清明节', '清明节.mp4', '清明节.jpg', 0, 0, 0, 1652597000, 1);
INSERT INTO `videos` VALUES (22, '春节', '春节.mp4', '春节.jpg', 0, 0, 0, 1652597010, 2);
INSERT INTO `videos` VALUES (23, '端午节', '端午节.mp4', '端午节.jpg', 0, 0, 0, 1652597020, 3);
INSERT INTO `videos` VALUES (24, '儿童节', '儿童节.mp4', '儿童节.jpg', 0, 0, 0, 1652597030, 4);
INSERT INTO `videos` VALUES (25, '护士节', '护士节.mp4', '护士节.jpg', 0, 0, 0, 1652597040, 5);
INSERT INTO `videos` VALUES (26, '圣诞节', '圣诞节.mp4', '圣诞节.jpg', 0, 0, 0, 1652597050, 6);
INSERT INTO `videos` VALUES (27, '万圣节', '万圣节.mp4', '万圣节.jpg', 0, 0, 0, 1652597060, 7);
INSERT INTO `videos` VALUES (28, '元旦节', '元旦节.mp4', '元旦节.jpg', 0, 0, 0, 1652597070, 1);
INSERT INTO `videos` VALUES (29, '元宵节', '元宵节.mp4', '元宵节.jpg', 0, 0, 0, 1652597080, 2);
INSERT INTO `videos` VALUES (30, '中秋节', '中秋节.mp4', '中秋节.jpg', 0, 0, 0, 1652597090, 3);
INSERT INTO `videos` VALUES (31, 'Forever missed', 'Forever missed.mp4', 'Forever missed.png', 0, 0, 0, 1652597090, 2);
INSERT INTO `videos` VALUES (32, 'fighting', 'fighting.mp4', 'fighting.png', 0, 0, 0, 1652597010, 2);
INSERT INTO `videos` VALUES (33, 'Hello comrades', 'Hello comrades.mp4', 'Hello comrades.png', 0, 0, 0, 1652597020, 3);
INSERT INTO `videos` VALUES (34, 'Long live Chairman Mao', 'Long live Chairman Mao.mp4', 'Long live Chairman Mao.png', 0, 0, 0, 1652597030, 4);
INSERT INTO `videos` VALUES (35, 'Long live the proletariat', 'Long live the proletariat.mp4', 'Long live the proletariat.png', 0, 0, 0, 1652597040, 5);
INSERT INTO `videos` VALUES (36, 'Remember the class struggle', 'Remember the class struggle.mp4', 'Remember the class struggle.png', 0, 0, 0, 1652597050, 6);
INSERT INTO `videos` VALUES (37, 'remember', 'remember.mp4', 'remember.png', 0, 0, 0, 1652597060, 7);
INSERT INTO `videos` VALUES (38, 'Serve the people wholeheartedly', 'Serve the people wholeheartedly.mp4', 'Serve the people wholeheartedly.png', 0, 0, 0, 1652597070, 1);
INSERT INTO `videos` VALUES (39, 'struggle', 'struggle.mp4', 'struggle.png', 0, 0, 0, 1652597080, 2);
INSERT INTO `videos` VALUES (40, 'Study well every day', 'Study well every day.mp4', 'Study well every day.png', 0, 0, 0, 1652597090, 3);
INSERT INTO `videos` VALUES (1000, '这是第一条视频的内容介绍：XXXX真可爱！', 'bear.mp4', 'bear.jpg', 10, 10, 0, 1652597777, 1);

SET FOREIGN_KEY_CHECKS = 1;
