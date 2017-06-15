-- phpMyAdmin SQL Dump
-- version 4.5.4.1deb2ubuntu2
-- http://www.phpmyadmin.net
--
-- Host: localhost
-- Generation Time: Jun 15, 2017 at 05:31 PM
-- Server version: 5.7.18-0ubuntu0.16.04.1-log
-- PHP Version: 7.0.18-0ubuntu0.16.04.1

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `qcrm`
--

-- --------------------------------------------------------

--
-- Table structure for table `clients`
--

CREATE TABLE `clients` (
  `id` int(10) UNSIGNED NOT NULL,
  `first_name` varchar(25) NOT NULL,
  `last_name` varchar(25) NOT NULL,
  `email` varchar(25) NOT NULL,
  `secret` varchar(200) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT NULL,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `clients`
--

INSERT INTO `clients` (`id`, `first_name`, `last_name`, `email`, `secret`, `created_at`, `updated_at`, `deleted_at`) VALUES
(12, 'khurshid', 'shah', 'ganesh@qwentic.com', '$2a$10$QIQh1Ix88KBdTBK98523m.AMPIpBRIL80KJiW5ae27VQr2Fejp/8q', '2017-06-14 06:24:35', '2017-06-14 06:24:35', NULL),
(13, 'khurshid', 'shah', 'khurshid@qwentic.com', '$2a$10$6UA2zcTYCbGgeJX46f.V0evlmRNwZst.D4B/2jHKlTWB1KHyhnUfa', '2017-06-14 06:25:18', '2017-06-14 06:25:18', NULL),
(16, 'deeksha', 'naik', 'deeksha@qwentic.com', '$2a$10$/7GKPskMhg9KaddTuNm4puj.Bfe.uIeCwKF87mSHYzlB31TcI.Ad6', '2017-06-14 06:38:42', '2017-06-14 06:38:42', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `companies`
--

CREATE TABLE `companies` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(500) NOT NULL,
  `responsible` varchar(250) NOT NULL,
  `type` varchar(25) NOT NULL,
  `industry_id` int(10) UNSIGNED NOT NULL,
  `employees` varchar(15) NOT NULL,
  `annual_income` float NOT NULL,
  `currency` varchar(25) NOT NULL,
  `comment` varchar(1000) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `companies`
--

INSERT INTO `companies` (`id`, `name`, `responsible`, `type`, `industry_id`, `employees`, `annual_income`, `currency`, `comment`, `created_at`, `updated_at`, `deleted_at`) VALUES
(2, 'Qwentic consulting pvt ltd', 'Pankaj', 'Client', 23, '0', 20, 'US Dollar', '', '2017-06-15 11:02:22', '2017-06-15 11:02:22', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `contacts`
--

CREATE TABLE `contacts` (
  `id` int(10) UNSIGNED NOT NULL,
  `first_name` varchar(25) NOT NULL,
  `last_name` varchar(25) NOT NULL,
  `image` text NOT NULL,
  `email` varchar(25) NOT NULL,
  `phone` varchar(12) NOT NULL,
  `company_id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `contacts`
--

INSERT INTO `contacts` (`id`, `first_name`, `last_name`, `image`, `email`, `phone`, `company_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(4, 'khurshid', 'shah', '', 'khurshid@qwentic.com', '1234567890', 2, '2017-06-15 12:00:20', '2017-06-15 12:00:20', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `industries`
--

CREATE TABLE `industries` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(50) NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `industries`
--

INSERT INTO `industries` (`id`, `name`, `created_at`, `updated_at`, `deleted_at`) VALUES
(23, 'Information Technology', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(24, 'Telecommunication', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(25, 'Manufacturing', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(26, 'Banking Services', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(27, 'Consulting', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(28, 'Finance', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(29, 'Government', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(30, 'Delivery', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(31, 'Non-Profit', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(32, 'Entertainment', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL),
(33, 'Other', '2017-06-15 10:16:07', '2017-06-15 10:16:07', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `sites`
--

CREATE TABLE `sites` (
  `id` int(10) UNSIGNED NOT NULL,
  `name` varchar(50) NOT NULL,
  `url` varchar(250) NOT NULL,
  `contact_id` int(10) UNSIGNED NOT NULL,
  `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  `updated_at` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
  `deleted_at` timestamp NULL DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `sites`
--

INSERT INTO `sites` (`id`, `name`, `url`, `contact_id`, `created_at`, `updated_at`, `deleted_at`) VALUES
(1, 'facebook', 'facebook.com', 4, '2017-06-15 12:00:20', '2017-06-15 12:00:20', NULL),
(2, 'linkedin', 'linkedin.com', 4, '2017-06-15 12:00:20', '2017-06-15 12:00:20', NULL);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `clients`
--
ALTER TABLE `clients`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `companies`
--
ALTER TABLE `companies`
  ADD PRIMARY KEY (`id`),
  ADD KEY `industry_id` (`industry_id`);

--
-- Indexes for table `contacts`
--
ALTER TABLE `contacts`
  ADD PRIMARY KEY (`id`),
  ADD KEY `company_id` (`company_id`),
  ADD KEY `company_id_2` (`company_id`);

--
-- Indexes for table `industries`
--
ALTER TABLE `industries`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `sites`
--
ALTER TABLE `sites`
  ADD PRIMARY KEY (`id`),
  ADD KEY `contact_id` (`contact_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `clients`
--
ALTER TABLE `clients`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=17;
--
-- AUTO_INCREMENT for table `companies`
--
ALTER TABLE `companies`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- AUTO_INCREMENT for table `contacts`
--
ALTER TABLE `contacts`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=5;
--
-- AUTO_INCREMENT for table `industries`
--
ALTER TABLE `industries`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=34;
--
-- AUTO_INCREMENT for table `sites`
--
ALTER TABLE `sites`
  MODIFY `id` int(10) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `companies`
--
ALTER TABLE `companies`
  ADD CONSTRAINT `industry_id` FOREIGN KEY (`industry_id`) REFERENCES `industries` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `contacts`
--
ALTER TABLE `contacts`
  ADD CONSTRAINT `company_id` FOREIGN KEY (`company_id`) REFERENCES `companies` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

--
-- Constraints for table `sites`
--
ALTER TABLE `sites`
  ADD CONSTRAINT `contact_id` FOREIGN KEY (`contact_id`) REFERENCES `contacts` (`id`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
