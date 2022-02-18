-- phpMyAdmin SQL Dump
-- version 5.1.1
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Waktu pembuatan: 19 Sep 2021 pada 13.55
-- Versi server: 10.4.20-MariaDB
-- Versi PHP: 7.4.22

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `db_books`
--

-- --------------------------------------------------------

--
-- Struktur dari tabel `books`
--

CREATE TABLE `books` (
  `id` int(11) NOT NULL,
  `title` varchar(255) NOT NULL,
  `description` varchar(500) NOT NULL,
  `image_url` varchar(500) NOT NULL,
  `release_year` int(11) NOT NULL,
  `price` varchar(255) NOT NULL,
  `total_page` varchar(255) NOT NULL,
  `kategori_ketebalan` varchar(255) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

--
-- Dumping data untuk tabel `books`
--

INSERT INTO `books` (`id`, `title`, `description`, `image_url`, `release_year`, `price`, `total_page`, `kategori_ketebalan`, `created_at`, `updated_at`) VALUES
(1, 'Algoritma', 'Buku tentang Algoritma', 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fwww.slideshare.net%2Farbo77%2Falgoritma-pemrograman-flowchart&psig=AOvVaw1NT5YYAbWnfC9d4qZI29_4&ust=1632136407297000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCNjAwoj0ivMCFQAAAAAdAAAAABAI', 2000, 'Rp. 150.000,-', '250', 'tebal', '2021-09-19 13:13:00', '2021-09-19 13:13:00'),
(2, 'Ganjil Genap', 'Ganjil Genap bukan buku tentang peraturan plat kendaraan di jalan protokol Jakarta, ya. Novel ini bercerita tentang kisah percintaan yang terjadi pada banyak orang. Bisa saja terjadi pada sahabat, atau bahkan kamu sendiri.\r\n\r\n', 'https://www.rukita.co/stories/wp-content/uploads/2020/01/buku-2020.jpg', 2020, 'Rp. 100.000,-', '144', 'sedang', '2021-09-19 13:17:00', '2021-09-19 13:17:00'),
(3, 'Sewu Dino', 'Buku horor yang dibuat oleh simpleman', 'https://www.rukita.co/stories/wp-content/uploads/2020/01/buku-2020-2.jpg', 2020, 'Rp. 65.000,-', '75', 'tipis', '2021-09-19 13:19:41', '2021-09-19 13:19:41'),
(4, 'Orang orang Biasa', 'Buku ciptaan Andrea Hirata', 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fmizanstore.com%2Forang-orang_biasa_64666&psig=AOvVaw3jqBgZBF-eBB_hEkQUKfO2&ust=1632136973127000&source=images&cd=vfe&ved=0CAwQjhxqFwoTCMDv5Jz2ivMCFQAAAAAdAAAAABAD', 2019, 'Rp. 80.000,-', '89', 'tipis', '2021-09-19 13:22:22', '2021-09-19 13:22:22'),
(5, 'To Kill The MockingBird', 'Buku yang terinspirasi dari kisah pembunuhan Richard Hickock dan Perry Smith pada tahun 1960 ini juga dinobatkan sebagai klasik sastra yang paling berpengaruh di dunia.', 'https://blog.mizanstore.com/wp-content/uploads/2019/09/To-Kill-A-Mockingbird-Edit-300x300.jpg', 1960, 'Rp. 125.000,-', '166', 'sedang', '2021-09-19 13:23:50', '2021-09-19 13:23:50'),
(6, 'Pride and Prejudice', 'Buku Pride and Prejudice adalah salah satu karya terlaris yang Jane Austen pernah tulis. Setidaknya, 20 juta eksemplar buku Pride and Prejudice yang telah diproduksi dan tersebar di seluruh dunia. Jane Austen menuangkan detail yang memikat mengenai kaum m', 'https://blog.mizanstore.com/wp-content/uploads/2019/09/Pride-Prejudice-Edit-300x300.jpg', 1813, 'Rp. 67.000,-', '88', 'tipis', '2021-09-19 13:25:46', '2021-09-19 13:25:46'),
(7, 'Gulag', 'Fiesta adalah buku peraih nobel sastra di tahun 1954, maka Gulag menjadi penerusnya. Novel ini telah meraih penghargaan nobel sastra di tahun 1970. Gulag bahkan disebut-sebut karya terbaik yang pernah ditulis oleh Aleksandr Isayevich Solzhenitsyn, sastrawan berbakat asal Rusia. Novel ini telah menjadi warisan sastra dunia dan telah diterjemahkan ke dalam berbagai bahasa, salah satunya, bahasa Indonesia.\r\n\r\n ', 'https://blog.mizanstore.com/wp-content/uploads/2019/09/Gulag-Edit-300x300.jpg', 1973, 'Rp. 89.000,-', '76', 'tipis', '2021-09-19 13:29:00', '2021-09-19 13:29:00'),
(8, 'What I Wish I Knew When I Was 20 oleh Teena Seelig', 'Dalam buku ini, penulisnya yaitu Tina Seelig membuat daftar hal-hal yang ia harap ia ketahui ketika akan kuliah untuk putranya yang akan kuliah. Ada banyak pesan dan nasihat yang dapat kamu ambil, seperti bagaimana menjadi lebih kreatif, menciptakan keberuntungan, serta nasihat untuk hidup dengan baik lainnya. ', 'https://suneducationgroup.com/wp-content/uploads/2020/10/What-I-Wish-I-Knew-When-I-Was-20.jpg', 2009, 'Rp. 156.000,-', '156', 'sedang', '2021-09-19 13:31:19', '2021-09-19 13:31:19'),
(9, 'How to Become a Straight-A Student', 'Karya ini merupakan tuntunan lengkap bagi kamu yang sama sekali belum memiliki gambaran mengenai kehidupan berkuliah. Kamu akan menemukan strategi agar kehidupanmu sebagai mahasiswa lebih terstruktur namun tetap menyenangkan untuk dijalani. Mulai dari strategi belajar, mempersiapkan ujian, membagi waktu, hingga bagaimana mengatasi menunda-nunda pekerjaan.', 'https://suneducationgroup.com/wp-content/uploads/2020/10/How-To-Become-A-Straight-A-Student.jpg', 2006, 'Rp. 60.000,-', '87', 'tipis', '2021-09-19 13:32:29', '2021-09-19 13:32:29'),
(10, 'The Alchemist', 'Buku ini adalah sumber dari ungkapan terkenal “when you really want something, the whole universe conspires in helping you to achieve it”. Novel yang menjadi secara umum memberikan gagasan mengenai “pencarian jati diri” yang menginspirasi pembacanya untuk mengejar apa yang diinginkan tanpa rasa takut. ', 'https://suneducationgroup.com/wp-content/uploads/2020/10/The-Alchemist.jpg', 1988, 'Rp. 142.000,-', '148', 'sedang', '2021-09-19 13:34:17', '2021-09-19 13:34:17'),
(11, 'Laskar Pelangi', 'Laskar Pelangi adalah novel pertama karya Andrea Hirata yang diterbitkan oleh Bentang Pustaka pada tahun 2005. Novel ini bercerita tentang kehidupan 10 anak dari keluarga miskin yang bersekolah (SD dan SMP) di sebuah sekolah Muhammadiyah di Belitung yang penuh dengan keterbatasan. Mereka adalah: Ikal aka Andrea Hirata.', 'https://upload.wikimedia.org/wikipedia/id/8/8e/Laskar_pelangi_sampul.jpg', 2008, 'Rp. 174.000,-', '534', 'tebal', '2021-09-19 13:35:32', '2021-09-19 13:35:32'),
(12, 'King SBMPTN 2020', 'Buku buatan dari KING untuk mengikuti ujian SBMPTN', 'https://images.tokopedia.net/img/cache/900/product-1/2020/4/18/11269917/11269917_7d16875c-8eb0-4dcd-b636-e070210f366d_1024_1024', 2020, 'Rp. 187.000,-', '823', 'tebal', '2021-09-19 13:38:35', '2021-09-19 13:38:35'),
(13, 'Wangsit SBMPTN 2020', 'Buku dengan judul pawang soal sulit atau WANGSIT', 'https://www.togamas.com/css/images/items/potrait/1123_11314_Wangsit__Pawang_Soal_Sulit_HOTS_SBMPTN_SAINTEK_2020__.jpg', 2020, 'Rp. 231.000,-', '857', 'tebal', '2021-09-19 13:41:10', '2021-09-19 13:41:10'),
(14, 'Panduan Praktis dan Lengkap Tahsin Tajwid Tahfizh', 'Buku yang memperbaiki bacaan Al-Quran', 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fdivapress-online.com%2Fbook%2Fpanduan-praktis-lengkap-tahsin-tajwid-tahfizh&psig=AOvVaw0SsoBDujnP1Mr7JnhEySmu&ust=1632138314164000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCIivrZb7ivMCFQAAAAAdAAAAABAD', 2019, 'Rp. 88.000,-', '224', 'tebal', '2021-09-19 13:44:20', '2021-09-19 13:44:20'),
(15, 'Cloud Computing', 'Buku tentang manajemen dan kapasitas cloud computing', 'https://www.google.com/url?sa=i&url=https%3A%2F%2Fshopee.co.id%2FCLOUD-COMPUTING-Manajemen-dan-Perencanaan-Kapasitas-Riko-Herwanto-Onno-W.-Purbo-RZ.-Abd.-Azis-i.156178717.4857899784&psig=AOvVaw0rCRtVBiVHnWHDszvRwgbW&ust=1632138540779000&source=images&cd=vfe&ved=0CAsQjRxqFwoTCIihtYT8ivMCFQAAAAAdAAAAABAD', 2020, 'Rp. 57.000,-', '168', 'sedang', '2021-09-19 13:46:17', '2021-09-19 13:46:17');

--
-- Indexes for dumped tables
--

--
-- Indeks untuk tabel `books`
--
ALTER TABLE `books`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT untuk tabel yang dibuang
--

--
-- AUTO_INCREMENT untuk tabel `books`
--
ALTER TABLE `books`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=16;
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
