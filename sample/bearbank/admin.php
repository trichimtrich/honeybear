<?php

if(isset($_GET["page"])) {
	if ($_GET["page"] === "staff") {
		if (isset($_GET["do"]) && $_GET["do"] === "add") {
			if (isset($_GET["step"]) && $_GET["step"] === "process")
				include "add_staff.php";
			else
				include "addstaff.php";
		}
		elseif (isset($_GET["do"]) && $_GET["do"] === "edit") {
			if (isset($_GET["step"]) && $_GET["step"] === "fill")
				include "editstaff.php";
			elseif (isset($_GET["step"]) && $_GET["step"] === "process")
				include "alter_staff.php";
			else
				include "display_staff.php";
		}
		elseif (isset($_GET["do"]) && $_GET["do"] === "delete") {
			if (isset($_GET["step"]) && $_GET["step"] === "process")
				include "editstaff.php";
			else
				include "delete_staff.php";
		}
		else
			include "display_staff.php";
	}

	elseif ($_GET["page"] === "customer") {
		if (isset($_GET["do"]) && $_GET["do"] === "add") {
			if (isset($_GET["step"]) && $_GET["step"] === "process")
				include "add_customer.php";
			else
				include "addcustomer.php";
		}
		elseif (isset($_GET["do"]) && $_GET["do"] === "edit") {
			if (isset($_GET["step"]) && $_GET["step"] === "fill")
				include "editcustomer.php";
			elseif (isset($_GET["step"]) && $_GET["step"] === "process")
				include "alter_customer.php";
			else
				include "display_customer.php";
		}
		elseif (isset($_GET["do"]) && $_GET["do"] === "delete") {
			if (isset($_GET["step"]) && $_GET["step"] === "process")
				include "editcustomer.php";
			else
				include "delete_customer.php";
		}
		else
			include "display_customer.php";
	}

	elseif ($_GET["page"] === "changepass")
		include "change_password.php";

	else
		include "admin_hompage.php";


} else {
	include "admin_hompage.php";
}

?>