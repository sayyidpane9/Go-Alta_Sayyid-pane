#!/bin/bash

#membuat folder utama dengan format nama "folder-name at date"
folderName="$1 at $(date)"
mkdir "$folder_Name"

#membuat subfolder about_me, about_me/personal, about_me/professional
mkdir "$folder_Name/about_me"
mkdir "$folder_Name/about_me/personal"
mkdir "$folder_Name/about_me/professional"

#membuat file facebook.txt dan linkedin.txt di subfolder about_me
echo "https://www.facebook.com/$2" > "$folder_Name/about_me/personal/facebook.txt"
echo "https://www.linkedin.com/in/$3" > "$folder_Name/about_me/professional/linkedin.txt"

#membuat subfolder my_friends
mkdir "$folder_Name/my_friends"

#mengambil daftar nama teman dari file list_of_my_friends.txt
curl -s "https://gist.githubusercontent.com/tegarimansyah/e91f335753ab2c7fb12815779677e914/raw/94864388379fecee450fde26e3e73bfb2bcda194/list%2520of%2520my%2520friends.txt" > "$folderName/my_friends/list_of_my_friends.txt"

#membuat subfolder my_system_info
mkdir "$folder_Name/my_system_info"
#menuliskan informasi tentang laptop ke file about_this_laptop.txt
echo "My username: $USER" > "$folder_Name/my_system_info/about_this_laptop.txt"
echo "With host: $(uname -a)" >> "$folder_Name/my_system_info/about_this_laptop.txt"

#menjalankan ping ke google.com sebanyak 3 kali dan menuliskan hasil ke file internet_connection.txt
ping -c 3 google.com > "$folder_Name/my_system_info/internet_connection.txt"

echo "Command automate telah selesai."

tree "$folder_Name"