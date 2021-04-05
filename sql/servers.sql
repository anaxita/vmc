CREATE TABLE IF NOT EXISTS `servers` (
  `id` varchar(255) PRIMARY KEY,
  `title` varchar(255) NOT NULL,
  `ip4` varchar(255) NOT NULL DEFAULT "",
  `hv` varchar(255) NOT NULL,
  `out_addr` varchar(255) NOT NULL DEFAULT "",
  `hostname` varchar(255) NOT NULL DEFAULT "",
  `description` varchar(255) NOT NULL DEFAULT "",
  `company` varchar(255) NOT NULL DEFAULT "",
  `user_name` varchar(255) NOT NULL DEFAULT "",
  `user_password` varchar(255) NOT NULL DEFAULT ""
);