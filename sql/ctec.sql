-- phpMyAdmin SQL Dump
-- version 5.2.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: May 19, 2024 at 11:10 PM
-- Server version: 10.4.32-MariaDB
-- PHP Version: 8.2.12

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `ctec`
--
CREATE DATABASE IF NOT EXISTS `ctec` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci;
USE `ctec`;

-- --------------------------------------------------------

--
-- Table structure for table `student_v2`
--

CREATE TABLE `student_v2` (
  `id` mediumint(8) UNSIGNED NOT NULL,
  `student_id` mediumint(9) NOT NULL,
  `first_name` varchar(255) NOT NULL,
  `last_name` varchar(255) NOT NULL,
  `email` varchar(255) NOT NULL,
  `phone` varchar(100) NOT NULL,
  `degree_program` varchar(255) NOT NULL,
  `gpa` float DEFAULT NULL,
  `financial_aid` tinyint(1) NOT NULL DEFAULT 0,
  `graduation_date` date DEFAULT NULL,
  `date_created` timestamp NOT NULL DEFAULT current_timestamp() ON UPDATE current_timestamp()
) ENGINE=InnoDB DEFAULT CHARSET=latin1 COLLATE=latin1_swedish_ci;

--
-- Dumping data for table `student_v2`
--

INSERT INTO `student_v2` (`id`, `student_id`, `first_name`, `last_name`, `email`, `phone`, `degree_program`, `gpa`, `financial_aid`, `graduation_date`, `date_created`) VALUES
(103, 1, 'Vittorio', 'Cornels', 'vcornels0@joomla.org', '883-429-8988', 'Cybersecurity', 0.51, 1, '2021-02-04', '2024-03-07 01:01:07'),
(104, 2, 'Dianna', 'Devennie', 'ddevennie1@ucoz.ru', '186-892-5962', 'Marketing', 0.46, 1, '2024-10-14', '2024-03-07 01:01:07'),
(105, 3, 'Amy', 'Stagg', 'astagg2@wp.com', '408-838-4973', 'Business Administration', NULL, 1, '2029-07-05', '2024-03-07 01:01:07'),
(106, 4, 'Ailbert', 'Atwell', 'aatwell3@cafepress.com', '297-924-9262', 'Professional Baking and Pastry Arts', 2.44, 1, '2026-12-23', '2024-03-07 01:01:07'),
(107, 5, 'Rona', 'Nancekivell', 'rnancekivell4@twitpic.com', '343-468-2121', 'Fine Arts', 2.22, 0, '2029-09-14', '2024-03-07 01:01:07'),
(108, 6, 'Stormy', 'Trainor', 'strainor5@ucoz.ru', '426-632-1979', 'Marketing', 0.94, 0, '2027-08-03', '2024-03-07 01:01:07'),
(109, 7, 'Camella', 'Sigars', 'csigars6@businessinsider.com', '745-243-7112', '', 0.91, 0, '2027-10-17', '2024-03-07 01:01:07'),
(110, 8, 'Boothe', 'Schiefersten', 'bschiefersten7@tmall.com', '274-867-6232', 'Marketing', 3.01, 1, '2023-01-19', '2024-03-07 01:01:07'),
(111, 9, 'Arnoldo', 'Ferrick', 'aferrick8@mysql.com', '307-219-8250', 'Advanced Manufacturing', 3.34, 0, '2029-06-02', '2024-03-07 01:01:07'),
(112, 10, 'Guthrie', 'Nicholls', 'gnicholls9@japanpost.jp', '764-667-8545', 'Cuisine Management', 0.9, 0, '2022-01-02', '2024-03-07 01:01:07'),
(113, 11, 'Sigismond', 'Sea', 'sseaa@shareasale.com', '337-688-0916', 'Cuisine Management', 1.86, 0, '2025-11-03', '2024-03-07 01:01:07'),
(114, 12, 'Alta', 'Andri', 'aandrib@mtv.com', '689-235-2628', 'Network Technology', 1.01, 0, '2022-03-18', '2024-03-07 01:01:07'),
(115, 13, 'Bellanca', 'Deinhard', 'bdeinhardc@answers.com', '976-727-7083', 'Cybersecurity', 0.93, 1, '2021-12-15', '2024-03-07 01:01:07'),
(116, 14, 'Jana', 'Issacov', 'jissacovd@mit.edu', '826-478-7355', 'Cybersecurity', NULL, 1, '2025-09-10', '2024-03-07 01:01:07'),
(117, 15, 'Abelard', 'Emslie', 'aemsliee@skyrock.com', '232-674-0514', 'Digital Media Arts', 0.08, 0, '2024-05-12', '2024-03-07 01:01:07'),
(118, 16, 'Astra', 'Buckston', 'abuckstonf@eventbrite.com', '504-472-3583', 'Network Technology', NULL, 0, '2024-06-07', '2024-03-07 01:01:07'),
(119, 17, 'Alvin', 'Balcock', 'abalcockg@dedecms.com', '726-431-6071', 'Digital Media Arts', 0.24, 1, '2024-10-13', '2024-03-07 01:01:07'),
(120, 18, 'Jourdain', 'Vollam', 'jvollamh@youtube.com', '680-793-1373', 'Cybersecurity', 1.63, 1, '2021-09-04', '2024-03-07 01:01:07'),
(121, 19, 'Glenda', 'Banfield', 'gbanfieldi@about.com', '750-321-7135', 'Business Administration', 2.35, 0, '2029-08-08', '2024-03-07 01:01:07'),
(122, 20, 'Cart', 'Cowthard', 'ccowthardj@51.la', '892-328-2563', 'Digital Media Arts', 1.06, 0, '2021-07-28', '2024-03-07 01:01:07'),
(123, 21, 'Bjorn', 'Frisdick', 'bfrisdickk@shutterfly.com', '165-550-7527', 'Advanced Manufacturing', 0.83, 1, NULL, '2024-03-07 01:01:07'),
(124, 22, 'Cassie', 'Cutress', 'ccutressl@go.com', '925-815-4864', 'Business Administration', 3.41, 0, '2026-03-28', '2024-03-07 01:01:07'),
(125, 23, 'Susi', 'Wattingham', 'swattinghamm@naver.com', '659-274-3964', 'Business Administration', 3.11, 0, '2021-04-30', '2024-03-07 01:01:07'),
(126, 24, 'Chrissy', 'McCusker', 'cmccuskern@nydailynews.com', '675-851-0330', 'Music', 2.71, 1, '2021-12-31', '2024-03-07 01:01:07'),
(127, 25, 'Joelie', 'Bickerton', 'jbickertono@digg.com', '260-997-8509', 'Music', 0.65, 0, '2028-02-03', '2024-03-07 01:01:07'),
(128, 26, 'Timothea', 'Ansteys', 'tansteysp@webnode.com', '501-755-9555', 'Cuisine Management', 1.49, 0, '2026-09-20', '2024-03-07 01:01:07'),
(129, 27, 'Colly', 'Tampling', 'ctamplingq@apache.org', '417-714-2880', 'Marketing', 3.51, 1, '2021-05-27', '2024-03-07 01:01:07'),
(130, 28, 'Marla', 'Kleuer', 'mkleuerr@amazonaws.com', '437-539-4192', '', NULL, 1, '2027-07-24', '2024-03-07 01:01:07'),
(131, 29, 'Vittoria', 'Sorel', 'vsorels@gnu.org', '537-248-9434', 'Fine Arts', 0.76, 0, '2029-08-08', '2024-03-07 01:01:07'),
(132, 30, 'Colet', 'Huddle', 'chuddlet@chron.com', '606-110-2737', 'Advanced Manufacturing', 1.57, 0, '2026-04-11', '2024-03-07 01:01:07'),
(133, 31, 'Kim', 'Dewfall', 'kdewfallu@amazon.co.jp', '756-735-8786', 'Web Development', 0.16, 1, '2022-04-28', '2024-03-07 01:01:07'),
(134, 32, 'Betteann', 'Bramall', 'bbramallv@google.com', '677-448-6607', 'Advanced Manufacturing', 0.87, 1, '2020-08-19', '2024-03-07 01:01:07'),
(135, 33, 'Dollie', 'Orpyne', 'dorpynew@storify.com', '821-877-2550', 'Music', 0.96, 1, '2028-12-18', '2024-03-07 01:01:07'),
(136, 34, 'Ulick', 'Guidera', 'uguiderax@google.pl', '377-609-8747', 'Professional Baking and Pastry Arts', NULL, 1, '2029-01-19', '2024-03-07 01:01:07'),
(137, 35, 'Phedra', 'Ickov', 'pickovy@wikia.com', '529-488-5043', 'Advanced Manufacturing', 0.83, 1, '2024-06-14', '2024-03-07 01:01:07'),
(138, 36, 'Cherilynn', 'Trevers', 'ctreversz@engadget.com', '220-114-0637', 'Digital Media Arts', 0.09, 0, '2026-05-19', '2024-03-07 01:01:07'),
(140, 38, 'Elvera', 'Lyddiatt', 'elyddiatt11@tinyurl.com', '186-211-1393', 'Professional Baking and Pastry Arts', NULL, 1, '2021-11-03', '2024-03-07 01:01:07'),
(141, 39, 'Wileen', 'Maxfield', 'wmaxfield12@prlog.org', '924-273-3937', 'Web Development', 3.83, 0, '2024-11-28', '2024-03-07 01:01:07'),
(142, 40, 'Stephen', 'Weller', 'sweller13@soup.io', '652-633-2956', 'Network Technology', NULL, 0, '2022-11-14', '2024-03-07 01:01:07'),
(143, 41, 'Allistir', 'Hawtrey', 'ahawtrey14@omniture.com', '394-714-0892', 'Digital Media Arts', 3.69, 1, '2027-08-13', '2024-03-07 01:01:07'),
(144, 42, 'Alfons', 'Boyet', 'aboyet15@123-reg.co.uk', '451-577-9606', 'Marketing', NULL, 0, '2027-04-21', '2024-03-07 01:01:07'),
(146, 44, 'Liane', 'Tash', 'ltash17@wufoo.com', '610-995-4287', 'Fine Arts', 3.3, 0, '2022-09-17', '2024-03-07 01:01:07'),
(147, 45, 'Boote', 'Saltsberger', 'bsaltsberger18@zimbio.com', '297-757-9172', 'Advanced Manufacturing', 0.38, 0, NULL, '2024-03-07 01:01:07'),
(148, 46, 'Therese', 'Orable', 'torable19@sciencedirect.com', '812-248-4814', 'Advanced Manufacturing', 1.54, 1, NULL, '2024-03-07 01:01:07'),
(149, 47, 'Anders', 'Baldini', 'abaldini1a@mail.ru', '751-935-6631', 'Web Development', 1.08, 1, '2023-09-17', '2024-03-07 01:01:07'),
(150, 48, 'Florian', 'Kinney', 'fkinney1b@hud.gov', '266-982-7065', '', 3.38, 0, '2020-12-01', '2024-03-07 01:01:07'),
(151, 49, 'Nevin', 'Antoszczyk', 'nantoszczyk1c@sbwire.com', '874-394-4314', '', 2.78, 1, '2026-03-04', '2024-03-07 01:01:07'),
(152, 50, 'Briney', 'Crowcum', 'bcrowcum1d@usa.gov', '334-778-0590', 'Cuisine Management', 0.27, 1, '2023-03-24', '2024-03-07 01:01:07'),
(153, 51, 'Keir', 'Mathon', 'kmathon1e@nps.gov', '274-348-1018', 'Music', 0.98, 0, '2028-12-24', '2024-03-07 01:01:07'),
(154, 52, 'Olly', 'Tansley', 'otansley1f@goodreads.com', '609-365-3319', 'Cuisine Management', 3.13, 0, '2027-04-09', '2024-03-07 01:01:07'),
(155, 53, 'Irwinn', 'Vanini', 'ivanini1g@samsung.com', '180-718-1816', 'Advanced Manufacturing', 0.35, 0, '2020-10-10', '2024-03-07 01:01:07'),
(156, 54, 'Fletch', 'Eldershaw', 'feldershaw1h@webeden.co.uk', '283-541-7907', 'Professional Baking and Pastry Arts', 3.89, 0, '2020-02-10', '2024-03-07 01:01:07'),
(157, 55, 'Monika', 'Gebbie', 'mgebbie1i@furl.net', '339-248-3728', 'Management', NULL, 0, '2023-10-02', '2024-03-07 01:01:07'),
(158, 56, 'Theodor', 'Bye', 'tbye1j@php.net', '277-418-7091', 'Marketing', 3.75, 0, '2023-12-17', '2024-03-07 01:01:07'),
(159, 57, 'Chevalier', 'Jankovsky', 'cjankovsky1k@e-recht24.de', '635-560-3568', 'Fine Arts', 0.38, 1, '2027-04-29', '2024-03-07 01:01:07'),
(160, 58, 'Erroll', 'Gleadhall', 'egleadhall1l@php.net', '545-429-5196', 'Management', 0.9, 1, '2020-08-31', '2024-03-07 01:01:07'),
(161, 59, 'Cherye', 'O Mannion', 'comannion1m@chronoengine.com', '859-455-0142', '', 0.76, 1, '2025-01-11', '2024-03-07 01:01:07'),
(162, 60, 'Sarena', 'Iczokvitz', 'siczokvitz1n@privacy.gov.au', '638-317-5886', 'Cybersecurity', 1.42, 1, NULL, '2024-03-07 01:01:07'),
(163, 61, 'Zaneta', 'Theml', 'ztheml1o@yale.edu', '176-787-6727', 'Network Technology', 3.08, 1, '2027-02-23', '2024-03-07 01:01:07'),
(164, 62, 'Cullen', 'Iglesias', 'ciglesias1p@salon.com', '787-188-0076', 'Business Administration', 2.27, 0, '2021-04-28', '2024-03-07 01:01:07'),
(165, 63, 'Cal', 'Pither', 'cpither1q@aol.com', '338-681-5225', 'Digital Media Arts', 2.28, 1, '2021-04-18', '2024-03-07 01:01:07'),
(166, 64, 'Breena', 'McGlynn', 'bmcglynn1r@slideshare.net', '784-844-5772', 'Music', NULL, 0, NULL, '2024-03-07 01:01:07'),
(167, 65, 'Laverna', 'Dumbrill', 'ldumbrill1s@pagesperso-orange.fr', '938-914-3184', 'Digital Media Arts', 0.18, 1, '2025-04-02', '2024-03-07 01:01:07'),
(168, 66, 'Karrie', 'de Glanville', 'kdeglanville1t@si.edu', '143-187-8743', 'Digital Media Arts', 0.43, 1, '2020-05-12', '2024-03-07 01:01:07'),
(169, 67, 'Cathryn', 'Polglaze', 'cpolglaze1u@typepad.com', '995-821-3786', 'Marketing', 0.96, 1, '2029-11-17', '2024-03-07 01:01:07'),
(170, 68, 'Jennica', 'Sherston', 'jsherston1v@cbc.ca', '965-984-2094', 'Professional Baking and Pastry Arts', 0.39, 0, '2021-10-10', '2024-03-07 01:01:07'),
(171, 69, 'Caryl', 'Hurn', 'churn1w@symantec.com', '977-687-7545', 'Cuisine Management', 2.13, 0, '2023-10-12', '2024-03-07 01:01:07'),
(172, 70, 'Ram', 'Brittoner', 'rbrittoner1x@trellian.com', '284-301-7901', 'Advanced Manufacturing', 0.15, 1, '2022-02-16', '2024-03-07 01:01:07'),
(173, 71, 'Archibald', 'Vasilyevski', 'avasilyevski1y@tuttocitta.it', '241-742-9107', 'Management', 2.38, 1, '2022-10-18', '2024-03-07 01:01:07'),
(174, 72, 'Kent', 'Shortell', 'kshortell1z@gravatar.com', '887-972-6380', 'Cybersecurity', 1.74, 0, '2024-07-24', '2024-03-07 01:01:07'),
(175, 73, 'Goldy', 'McLernon', 'gmclernon20@wikimedia.org', '206-336-8310', 'Cybersecurity', 2.97, 1, '2024-09-17', '2024-03-07 01:01:07'),
(176, 74, 'Timmy', 'Eirwin', 'teirwin21@google.nl', '349-409-0568', 'Marketing', NULL, 1, '2026-05-22', '2024-03-07 01:01:07'),
(177, 75, 'Karalee', 'Probbin', 'kprobbin22@microsoft.com', '302-186-5842', 'Fine Arts', 2.39, 1, '2027-07-02', '2024-03-07 01:01:07'),
(178, 76, 'Jeana', 'Brikner', 'jbrikner23@shareasale.com', '746-740-8393', 'Network Technology', 2.58, 0, '2023-05-17', '2024-03-07 01:01:07'),
(179, 77, 'Emiline', 'Amies', 'eamies24@imgur.com', '498-407-2429', 'Marketing', 2.88, 0, '2025-09-10', '2024-03-07 01:01:07'),
(180, 78, 'Nolie', 'Gallihaulk', 'ngallihaulk25@dailymotion.com', '693-364-7932', 'Web Development', 2.55, 0, '2023-01-18', '2024-03-07 01:01:07'),
(181, 79, 'Matty', 'Teodori', 'mteodori26@mashable.com', '381-194-5838', 'Music', 3.91, 1, '2029-11-11', '2024-03-07 01:01:07'),
(182, 80, 'Devin', 'Gatiss', 'dgatiss27@w3.org', '855-155-7125', 'Fine Arts', 2.62, 0, '2021-08-03', '2024-03-07 01:01:07'),
(183, 81, 'Farlie', 'Agett', 'fagett28@europa.eu', '150-994-8613', 'Music', 2.47, 1, '2022-11-24', '2024-05-19 07:14:59'),
(184, 82, 'Eldon', 'Drayson', 'edrayson29@shop-pro.jp', '286-753-1377', '', 2.98, 1, '2028-07-07', '2024-03-07 01:01:07'),
(185, 83, 'Gabrila', 'Crowder', 'gcrowder2a@ibm.com', '619-856-4140', 'Fine Arts', 2.63, 0, '2022-12-25', '2024-03-07 01:01:07'),
(186, 84, 'Batholomew', 'Catenot', 'bcatenot2b@usda.gov', '621-483-3820', 'Marketing', NULL, 0, '2020-07-30', '2024-03-07 01:01:07'),
(187, 85, 'Terrill', 'McPhail', 'tmcphail2c@tumblr.com', '766-624-2461', 'Cybersecurity', 3.17, 1, '2020-03-07', '2024-03-07 01:01:07'),
(188, 86, 'Carolyne', 'Kelner', 'ckelner2d@cnet.com', '972-900-7045', 'Business Administration', 1.83, 1, '2020-04-18', '2024-03-07 01:01:07'),
(189, 87, 'Byram', 'Winsiowiecki', 'bwinsiowiecki2e@facebook.com', '847-951-9238', 'Cuisine Management', 3.25, 1, '2024-11-19', '2024-03-07 01:01:07'),
(190, 88, 'Aloise', 'Swynfen', 'aswynfen2f@joomla.org', '162-584-1074', 'Network Technology', 3.13, 1, '2027-04-02', '2024-03-07 01:01:07'),
(191, 89, 'Sharona', 'Cristofolini', 'scristofolini2g@umich.edu', '283-262-9213', 'Advanced Manufacturing', 3.01, 1, '2021-11-15', '2024-03-07 01:01:07'),
(192, 90, 'Almeta', 'Portam', 'aportam2h@dedecms.com', '190-277-5541', 'Digital Media Arts', 3.79, 0, '2022-11-05', '2024-03-07 01:01:07'),
(193, 91, 'Jacinthe', 'Faughny', 'jfaughny2i@com.com', '322-940-0373', 'Business Administration', 0.87, 1, '2020-07-20', '2024-03-07 01:01:07'),
(194, 92, 'Garrett', 'Dikes', 'gdikes2j@spotify.com', '188-510-4030', 'Music', NULL, 1, NULL, '2024-03-07 01:01:07'),
(195, 93, 'Muire', 'Edlyn', 'medlyn2k@boston.com', '449-974-7100', 'Management', 1.72, 1, NULL, '2024-03-07 01:01:07'),
(196, 94, 'Peggi', 'Winterbottom', 'pwinterbottom2l@1688.com', '831-440-0177', 'Cuisine Management', 1.6, 1, '2029-04-16', '2024-03-07 01:01:07'),
(197, 95, 'Kirbie', 'Mattecot', 'kmattecot2m@cdc.gov', '535-437-8747', 'Digital Media Arts', 0.43, 1, '2022-07-06', '2024-03-07 01:01:07'),
(198, 96, 'Alane', 'Angier', 'aangier2n@irs.gov', '137-183-3771', 'Management', 3.5, 1, '2025-09-04', '2024-03-07 01:01:07'),
(199, 97, 'Pate', 'Losel', 'plosel2o@vk.com', '284-402-7743', 'Business Administration', 3.73, 0, '2026-12-10', '2024-03-07 01:01:07'),
(200, 98, 'Pascale', 'Philippou', 'pphilippou2p@usa.gov', '965-957-3827', 'Digital Media Arts', 0.39, 1, '2020-05-28', '2024-03-07 01:01:07'),
(201, 99, 'Dall', 'McCadden', 'dmccadden2q@hud.gov', '422-567-3098', 'Network Technology', 2.04, 0, '2029-05-28', '2024-03-07 01:01:07'),
(202, 100, 'Dacey', 'Currey', 'dcurrey2r@rakuten.co.jp', '154-437-7551', 'Network Technology', 1.38, 1, '2028-01-30', '2024-03-07 01:01:07');

--
-- Indexes for dumped tables
--

--
-- Indexes for table `student_v2`
--
ALTER TABLE `student_v2`
  ADD PRIMARY KEY (`id`),
  ADD UNIQUE KEY `student_id` (`student_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `student_v2`
--
ALTER TABLE `student_v2`
  MODIFY `id` mediumint(8) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=229;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
