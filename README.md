# Shape-API

I used Gin Framework, JWT and Postgresql to build Login/Register features
Database:
<img width="825" alt="image" src="https://user-images.githubusercontent.com/39989729/179405889-f829a25a-998f-4b01-b907-e68fcfc00f67.png">

To reduce delay time every request area or perimeter so I used Redis to cache shape info include area and perimeter.
When user create or update shape. Service will calculate and save to redis again.

This image below show how it work:
<img width="1192" alt="image" src="https://user-images.githubusercontent.com/39989729/179392107-14303823-33d4-48ba-86ae-ce08b4e436f1.png">
