-- phpMyAdmin SQL Dump
-- version 5.0.1
-- https://www.phpmyadmin.net/
--
-- 主机： 127.0.0.1
-- 生成日期： 2020-05-18 05:22:33
-- 服务器版本： 10.4.11-MariaDB
-- PHP 版本： 7.4.3

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- 数据库： `gorm`
--

-- --------------------------------------------------------

--
-- 表的结构 `address`
--

CREATE TABLE `address` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) UNSIGNED NOT NULL,
  `addr` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `address`
--

INSERT INTO `address` (`id`, `user_id`, `addr`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 8, '安徽合肥市蜀山区110号', '2020-05-07 06:36:36', '2020-05-07 06:36:36', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `credit_card`
--

CREATE TABLE `credit_card` (
  `id` int(11) UNSIGNED NOT NULL,
  `user_id` int(11) UNSIGNED NOT NULL,
  `no` varchar(255) NOT NULL,
  `issue` varchar(266) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `credit_card`
--

INSERT INTO `credit_card` (`id`, `user_id`, `no`, `issue`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 10, '62234324234324', '中国工商银行', '2020-05-07 07:08:25', '2020-05-07 07:08:25', NULL),
(2, 11, '62234324234324', '中国工商银行', '2020-05-07 07:26:04', '2020-05-07 07:26:04', NULL),
(3, 11, '46867876867867', '中国建设银行', '2020-05-07 07:26:04', '2020-05-07 07:26:04', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `language`
--

CREATE TABLE `language` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT current_timestamp(),
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `language`
--

INSERT INTO `language` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, '中文', '2020-05-07 08:17:11', '2020-05-07 08:38:14', NULL),
(2, '英文', '2020-05-07 08:17:11', '2020-05-07 08:34:21', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `user`
--

CREATE TABLE `user` (
  `id` int(11) UNSIGNED NOT NULL,
  `name` varchar(255) NOT NULL,
  `age` tinyint(2) UNSIGNED NOT NULL,
  `birthday` date DEFAULT NULL,
  `created_at` timestamp NULL DEFAULT NULL,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` datetime DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `user`
--

INSERT INTO `user` (`id`, `name`, `age`, `birthday`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'Jinzhu', 10, '1990-10-10', '2020-05-07 02:13:59', '2020-05-07 02:13:59', NULL),
(2, 'Jinzhu', 11, '1990-10-10', '2020-05-07 02:14:03', '2020-05-07 02:14:03', NULL),
(3, 'Tim', 12, '1990-10-10', '2020-05-07 02:14:06', '2020-05-07 02:14:06', NULL),
(4, 'Jack', 13, '1990-10-10', '2020-05-07 02:14:09', '2020-05-07 02:14:09', NULL),
(5, 'Jinzhu', 14, '1990-10-10', '2020-05-07 02:14:12', '2020-05-07 02:14:12', NULL),
(6, 'Yoyo', 15, '1990-10-10', '2020-05-07 02:14:15', '2020-05-07 02:14:15', NULL),
(8, '小红', 18, '2000-10-10', '2020-05-07 06:36:36', '2020-05-07 06:36:36', NULL),
(10, '小明', 22, '1998-12-12', '2020-05-07 07:08:24', '2020-05-07 07:08:24', NULL),
(11, '小花', 22, '1998-12-12', '2020-05-07 07:26:04', '2020-05-07 07:26:04', NULL),
(13, '小天', 32, '1988-12-12', '2020-05-07 08:17:11', '2020-05-07 08:17:11', NULL),
(19, '小喜', 32, '1988-12-12', '2020-05-07 08:34:21', '2020-05-07 08:34:21', NULL),
(20, '小欢', 28, '1992-12-12', '2020-05-07 08:38:14', '2020-05-07 08:38:14', NULL),
(21, 'Giraffe', 0, '0000-00-00', '2020-05-08 09:16:05', '2020-05-08 09:16:05', NULL),
(22, 'Lion', 0, '0000-00-00', '2020-05-08 09:16:05', '2020-05-08 09:16:05', NULL);

-- --------------------------------------------------------

--
-- 表的结构 `user_languages`
--

CREATE TABLE `user_languages` (
  `user_id` int(11) UNSIGNED NOT NULL,
  `language_id` int(11) UNSIGNED NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- 转存表中的数据 `user_languages`
--

INSERT INTO `user_languages` (`user_id`, `language_id`) VALUES
(13, 1),
(13, 2),
(19, 1),
(19, 2),
(20, 1);

--
-- 转储表的索引
--

--
-- 表的索引 `address`
--
ALTER TABLE `address`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- 表的索引 `credit_card`
--
ALTER TABLE `credit_card`
  ADD PRIMARY KEY (`id`),
  ADD KEY `user_id` (`user_id`);

--
-- 表的索引 `language`
--
ALTER TABLE `language`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `name` (`name`);

--
-- 表的索引 `user`
--
ALTER TABLE `user`
  ADD PRIMARY KEY (`id`);

--
-- 表的索引 `user_languages`
--
ALTER TABLE `user_languages`
  ADD KEY `user_id` (`user_id`),
  ADD KEY `language_id` (`language_id`);

--
-- 在导出的表使用AUTO_INCREMENT
--

--
-- 使用表AUTO_INCREMENT `address`
--
ALTER TABLE `address`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- 使用表AUTO_INCREMENT `credit_card`
--
ALTER TABLE `credit_card`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- 使用表AUTO_INCREMENT `language`
--
ALTER TABLE `language`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;

--
-- 使用表AUTO_INCREMENT `user`
--
ALTER TABLE `user`
  MODIFY `id` int(11) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=23;

--
-- 限制导出的表
--

--
-- 限制表 `address`
--
ALTER TABLE `address`
  ADD CONSTRAINT `address_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- 限制表 `credit_card`
--
ALTER TABLE `credit_card`
  ADD CONSTRAINT `credit_card_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- 限制表 `user_languages`
--
ALTER TABLE `user_languages`
  ADD CONSTRAINT `user_languages_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `user` (`id`) ON DELETE CASCADE ON UPDATE CASCADE,
  ADD CONSTRAINT `user_languages_ibfk_2` FOREIGN KEY (`language_id`) REFERENCES `language` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
