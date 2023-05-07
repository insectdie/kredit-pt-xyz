-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: May 07, 2023 at 06:57 PM
-- Server version: 10.4.28-MariaDB
-- PHP Version: 8.1.17

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `kredit_pt_xyz`
--

-- --------------------------------------------------------

--
-- Table structure for table `konsumen`
--

CREATE TABLE `konsumen` (
  `nik` varchar(16) NOT NULL,
  `full_name` varchar(60) NOT NULL,
  `legal_name` varchar(60) NOT NULL,
  `tempat_lahir` varchar(30) NOT NULL,
  `tanggal_lahir` date NOT NULL,
  `gaji` int(11) NOT NULL,
  `foto_ktp` varchar(255) NOT NULL,
  `foto_selfie` varchar(255) NOT NULL,
  `created_datetime` datetime DEFAULT current_timestamp(),
  `modified_datetime` datetime DEFAULT NULL ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `konsumen`
--

INSERT INTO `konsumen` (`nik`, `full_name`, `legal_name`, `tempat_lahir`, `tanggal_lahir`, `gaji`, `foto_ktp`, `foto_selfie`, `created_datetime`, `modified_datetime`) VALUES
('1111111111111111', 'BUDI', 'BUDI', 'MALANG', '2000-10-10', 4000000, 'qadaszxczx.jpg', 'qadassadzxczx.jpg', '2023-05-06 19:22:33', '2023-05-07 20:15:17'),
('1234567890123456', 'ANNISA', 'ANNISA', 'PEKAN BARU', '1999-02-01', 10000000, 'asdkasjhdkashjd.jpg', 'asdkasjhdasdsakashjd.jpg', '2023-05-07 20:14:56', NULL);

-- --------------------------------------------------------

--
-- Table structure for table `konsumen_limit`
--

CREATE TABLE `konsumen_limit` (
  `nik` varchar(16) NOT NULL,
  `bulan` int(11) NOT NULL,
  `limit_pinjaman` int(11) NOT NULL,
  `created_datetime` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Dumping data for table `konsumen_limit`
--

INSERT INTO `konsumen_limit` (`nik`, `bulan`, `limit_pinjaman`, `created_datetime`) VALUES
('1111111111111111', 1, 100000, '2023-05-07 20:16:31'),
('1111111111111111', 2, 200000, '2023-05-07 20:17:06'),
('1111111111111111', 3, 500000, '2023-05-07 20:17:34'),
('1111111111111111', 4, 700000, '2023-05-07 20:17:57'),
('1234567890123456', 1, 1000000, '2023-05-07 20:16:31'),
('1234567890123456', 2, 1200000, '2023-05-07 20:17:06'),
('1234567890123456', 3, 1500000, '2023-05-07 20:17:34'),
('1234567890123456', 4, 2000000, '2023-05-07 20:17:57');

-- --------------------------------------------------------

--
-- Table structure for table `transaksi`
--

CREATE TABLE `transaksi` (
  `no_kontrak` int(7) NOT NULL,
  `nik` varchar(16) NOT NULL,
  `otr` int(11) NOT NULL,
  `admin_fee` int(11) NOT NULL,
  `jml_cicilan` int(11) NOT NULL,
  `jml_bunga` int(11) NOT NULL,
  `nama_asset` varchar(255) NOT NULL,
  `created_datetime` datetime DEFAULT current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `konsumen`
--
ALTER TABLE `konsumen`
  ADD PRIMARY KEY (`nik`);

--
-- Indexes for table `konsumen_limit`
--
ALTER TABLE `konsumen_limit`
  ADD PRIMARY KEY (`nik`,`bulan`);

--
-- Indexes for table `transaksi`
--
ALTER TABLE `transaksi`
  ADD PRIMARY KEY (`no_kontrak`,`nik`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `transaksi`
--
ALTER TABLE `transaksi`
  MODIFY `no_kontrak` int(7) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=1111129;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
