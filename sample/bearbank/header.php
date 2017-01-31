
    <body>
        <div class="wrapper">
        <div class="header">
        <?php
        if ($_SERVER["SERVER_ADDR"]=="192.168.200.107") { ?>
            <img src="header.png" height="300" width="100%"/>
        <?php } else { ?>
            <img src="header_hacker.png" height="300" width="100%"/>
        <?php } ?>
            </div>
            <div class="navbar">
                
            <ul>
            <li><a href="index.php">Home</a></li>
            <li><a href="features.php">Features</a></li>
            <li id="last"><a href="contact.php">Contact Us</a></li>
            </ul>
            </div>