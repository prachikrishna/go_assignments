package main

import (
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"
)

func main() {
	s := `
<!DOCTYPE html> 
<html lang="en-US"> 
<head>
<title>Biggest Online Tutorials Library</title>
<meta charset="utf-8">
<meta http-equiv="X-UA-Compatible" content="IE=edge">
<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1">
<meta name="Description" content="Biggest Online Tutorials Library - The Best Content on latest technologies including C, C++, Java, Python, PHP, Machine Learning, Data Science, AppML, AI with Python, Behave, Java16, Spacy."/>
<meta name="Keywords" content="HTML, Python, CSS, SQL, JavaScript, How to, PHP, Java, C++, jQuery, Bootstrap, C#, XML, MySQL, NodeJS, React, Angular, R, AI, Git, Data Science, Code Game, Tutorials, Programming, Web Development, Training, Learning, Quiz, Courses, Lessons, References, Examples, Source code, Demos, Tips"/>
<link rel="canonical" href="https://www.tutorialspoint.com" />
<link rel="stylesheet" href="/themes/home/tp-home.css?v=1.51" />
</head>	
<body>
<!-- top banner content -->	
<div class="tp-search-widget"> 
	<div class="mui-appbar-home">
		<div class="mui-container">
		   <ul>
			   <li><a href="https://www.tutorialspoint.com/index.htm" target="_blank" title="TutorialsPoint - Home"><i class="fa fa-home"></i> <span>Home</span></a></li>
			   <li><a href="https://www.tutorialspoint.com/about/about_careers.htm" target="_blank" title="Job @ Tutorials Point"><i class="fa fa-suitcase"></i> <span>Jobs</span></a></li>
			   <li><a href="https://www.tutorialspoint.com/questions/index.php" target="_blank" title="Questions &amp; Answers - The Best Technical Questions and Answers - TutorialsPoint"><i class="fa fa-location-arrow"></i> <span>Q/A</span></a></li>
			   <li><a href="https://www.tutorialspoint.com/online_dev_tools.htm" target="_blank" title="Tools - Online Development and Testing Tools"><i class="fa fa-cogs"></i> <span>Tools</span></a></li>
			   <li><a href="https://www.tutorialspoint.com/codingground.htm" target="_blank" title="Coding Ground - Free Online IDE and Terminal"><i class="fa fa-code"></i> <span>Coding Ground </span></a></li>
			   <li><a href="https://www.tutorialspoint.com/upsc_ias_exams.htm" target="_blank" title="UPSC IAS Exams Notes - TutorialsPoint"><i class="fa fa-user-tie"></i> <span>UPSC Notes</span></a></li>
			   <li><a href="https://www.tutorialspoint.com/whiteboard.htm" target="_blank" title="Free Online Whiteboard"><i class="fa fa-chalkboard"></i> <span>Whiteboard</span></a></li>
			   <li><a href="https://www.tutorix.com" target="_blank" title="Tutorx - The Best Learning App" rel="nofollow"><b><i class="fa fa-video"></i></b> <span>Tutorix</span></a></li>				
			   <li class="appbar-right"><a href="https://www.tutorialspoint.com/market/login.asp?v=4.00" title="Tutorialspoint" rel="nofollow"><b><i class="fa fa-sign-in-alt"></i> </b> <span>Login</span></a></li>
			   <li class="appbar-right"><a href="https://www.tutorialspoint.com/market/signup.asp?v=4.00" title="Tutorialspoint" rel="nofollow"><b><i class="fa fa-sign-out"></i> </b> <span>Sign up</span></a></li>
		   </ul>			
		</div>
		<div class="clear"></div>
	</div>
	<div class="mui-container">
		<div class="tp-logo">
			<a href="https://www.tutorialspoint.com/index.htm" title="Tutorialspoint"></a>
		</div>
	   <h4>You are browsing the best resource for <b>Online Education</b></h4>
	   <div class="search-box" style="display: block;">
	   <div>
	   <input class="input" type="text" id="search-strings" name="key" placeholder="Search latest courses, ebooks and prime packs..." autocomplete="off">
	   <button id="btnSearch"><i class="fa fa-search" title="Search"></i></button>
	   </div>
	   </div>		
	   <div class="search-box" id="search-results"></div>
	   <div class="clear"></div>
	   <div class="tp-banner-links">
	   <a href="https://www.tutorialspoint.com/tutorialslibrary.htm"><i class="fa fa-books"></i> <span>Library</span></a> 
	   <a href="https://www.tutorialspoint.com/market/index.asp"><i class="fa fa-video"></i> <span>Courses</span></a>
	   <a href="https://www.tutorialspoint.com/latest/ebooks"><i class="fa fa-book-reader"></i> <span>eBooks</span></a>
	   <a href="https://www.tutorialspoint.com/latest/prime-packs"><i class="fa fa-cubes"></i> <span>Prime Packs</span></a>
	   </div>  
	</div>  	
	<div class="clear"></div>
</div> 
<!-- end of top banner content -->
	<div class="mui-container">
	      <a href="https://www.tutorialspoint.com/codingground.htm" title="Coding Ground"></a>
       </div>

<!-- video courses, ebooks and prime pack content starts here -->
<div class="mui-container videos-wrap np">
	<h1 class="uk-heading-line"><span><i class="fa fa-cubes"></i> Prime <b>Packs</b></span></h1>
	<div class="slides">
	<div class="mui-col-xl-3 mui-col-md-3 mui-col-sm-6 mui-col-xs-12">
		<div class="prime-pack rounded">
			<div class="ribbon-B"><span>Prime Pack</span></div>
			<div class="course-card-thumbnail rounded">
				<a href="https://www.tutorialspoint.com/prime-pack/fullstack-web-development-prime-pack/index.asp" target="_blank"><img class="rounded-img" src="https://d3mxt5v3yxgcsr.cloudfront.net/courses/4982/course_4982_image.jpg" alt="FullStack Web Development Prime Pack" title="FullStack Web Development Prime Pack">
				<span class="prime-icon-trigger"></span></a>
			</div>
			<div class="primepack-card-body">
				<h4 class="h48"><a href="https://www.tutorialspoint.com/prime-pack/fullstack-web-development-prime-pack/index.asp" title="FullStack Web Development Prime Pack">FullStack Web Development Prime Pack</a></h4>
				<p class="videos-details"><i class="fa fa-play-circle"></i> &nbsp;9 Courses &nbsp;&nbsp;
					<span><i class="fa fa-file-pdf"></i> &nbsp;2 eBooks </span>
				</p>
				<p class="h22">Created by - <a href="https://www.tutorialspoint.com/profile/tutorialspoint_com" target="_blank">Tutorialspoint</a></p>
				<div class="ebk-course-price">
					<div class="ebook-price" style="display: block;">
						<div class="">
							<b><span class="price" data-jpy="7999.99" data-usd="59.99" data-inr="2599.99" data-cad="79.99" data-aud="79.99" data-gbp="59.99" data-eur="59.99">₹2,599.99</span></b>
							<span class="actual-price" data-jpy="54599" data-usd="499.99" data-inr="36599.99" data-cad="609.99" data-aud="639.99" data-gbp="359.99" data-eur="414.99">₹36,599.99</span>
						</div>						
						<div class="ebook-discount">				
							<span>You Save</span> 
							<b><span class="discount-amount" data-jpy="46599.01" data-usd="440" data-inr="34000" data-cad="530" data-aud="560" data-gbp="300" data-eur="355">₹34,000.00</span></b>
							<span class="ebook-off rounded" data-jpy="85" data-usd="88" data-inr="93" data-cad="87" data-aud="88" data-gbp="83" data-eur="86" style="display: block;">93 % off</span> 		
						</div>
					</div>
					<div class="mui-col-md-12 mui-col-sm-12 add_to_cart_div np">
						<div class="row">
							<div class="btn-success rounded">
								<a href="https://www.tutorialspoint.com/prime-pack/fullstack-web-development-prime-pack/index.asp" target="_blank">Buy Now</a>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<div class="mui-col-xl-3 mui-col-md-3 mui-col-sm-6 mui-col-xs-12">
		<div class="prime-pack rounded">
			<div class="ribbon-B"><span>Prime Pack</span></div>
			<div class="course-card-thumbnail">
				<a href="https://www.tutorialspoint.com/prime-pack/mysql-technologies-pack/index.asp"><img class="rounded-img" src="https://d3mxt5v3yxgcsr.cloudfront.net/courses/4938/course_4938_image.jpg" alt="MySQL Technologies Pack" title="MySQL Technologies Pack">
				<span class="prime-icon-trigger"></span></a>
			</div>
			<div class="primepack-card-body">
				<h4 class="h48"><a href="https://www.tutorialspoint.com/prime-pack/mysql-technologies-pack/index.asp" title="MySQL Technologies Pack">MySQL Technologies Pack</a></h4>
				<p class="videos-details"><i class="fa fa-play-circle"></i> &nbsp;6 Courses &nbsp;&nbsp;
					<span><i class="fa fa-file-pdf"></i> &nbsp;1 eBooks </span>
				</p>
				<p class="h22">Created by - <a href="https://www.tutorialspoint.com/profile/tutorialspoint_com" target="_blank">Tutorialspoint</a></p>
				<div class="ebk-course-price">
					<div class="ebook-price" style="display: block;">
						<div class="">
							<b><span class="price" data-usd="44.99" data-inr="1799.99" data-cad="59.99" data-aud="59.99" data-gbp="44.99" data-eur="44.99" data-jpy="5999.99">₹1,799.99</span></b>
							<span class="actual-price" data-usd="399.99" data-inr="19999.99" data-cad="489.99" data-aud="510.99" data-gbp="289.99" data-eur="330.99" data-jpy="41999.99">₹19,999.99</span>
						</div>						
						<div class="ebook-discount">				
							<span>You Save</span> 
							<b><span class="discount-amount" data-usd="355" data-inr="18200" data-cad="430" data-aud="451" data-gbp="245" data-eur="286" data-jpy="36000">₹18,200.00</span></b>
							<span class="ebook-off rounded" data-usd="89" data-inr="91" data-cad="88" data-aud="88" data-gbp="84" data-eur="86" data-jpy="86" style="display: block;">91 % off</span> 		
						</div>
					</div>
					<div class="mui-col-md-12 mui-col-sm-12 add_to_cart_div np">
						<div class="row">
							<div class="btn-success rounded">
								<a href="https://www.tutorialspoint.com/prime-pack/mysql-technologies-pack/index.asp" target="_blank">Buy Now</a>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<div class="mui-col-xl-3 mui-col-md-3 mui-col-sm-6 mui-col-xs-12">
		<div class="prime-pack rounded">
			<div class="ribbon-B"><span>Prime Pack</span></div>
			<div class="course-card-thumbnail">
				<a href="https://www.tutorialspoint.com/prime-pack/complete-python-prime-pack/index.asp"><img class="rounded-img" src="https://d3mxt5v3yxgcsr.cloudfront.net/courses/4870/course_4870_image.jpg" alt="Complete Python Prime Pack" title="Complete Python Prime Pack">
				<span class="prime-icon-trigger"></span></a>
			</div>
			<div class="primepack-card-body">
				<h4 class="h48"><a href="https://www.tutorialspoint.com/prime-pack/complete-python-prime-pack/index.asp" title="Complete Python Prime Pack">Complete Python Prime Pack</a></h4>
				<p class="videos-details"><i class="fa fa-play-circle"></i> &nbsp;9 Courses &nbsp;&nbsp;
					<span><i class="fa fa-file-pdf"></i> &nbsp;2 eBooks </span>
				</p>
				<p class="h22">Created by - <a href="https://www.tutorialspoint.com/profile/tutorialspoint_com" target="_blank">Tutorialspoint</a></p>
				<div class="ebk-course-price">
					<div class="ebook-price" style="display: block;">
						<div class="">
							<b><span class="price" data-usd="59.99" data-inr="2599.99" data-cad="79.99" data-aud="79.99" data-gbp="59.99" data-eur="59.99" data-jpy="7999.99">₹2,599.99</span></b>
							<span class="actual-price" data-usd="499.99" data-inr="36684.27" data-cad="609.99" data-aud="639.99" data-gbp="359.99" data-eur="414.99" data-jpy="54358.99">₹36,684.27</span>
						</div>						
						<div class="ebook-discount">				
							<span>You Save</span> 
							<b><span class="discount-amount" data-usd="440" data-inr="34084.28" data-cad="530" data-aud="560" data-gbp="300" data-eur="355" data-jpy="46359">₹34,084.28</span></b>
							<span class="ebook-off rounded" data-usd="88" data-inr="93" data-cad="87" data-aud="88" data-gbp="83" data-eur="86" data-jpy="85" style="display: block;">93 % off</span> 		
						</div>
					</div>
					<div class="mui-col-md-12 mui-col-sm-12 add_to_cart_div np">
						<div class="row">
							<div class="btn-success rounded">
								<a href="https://www.tutorialspoint.com/prime-pack/complete-python-prime-pack/index.asp" target="_blank">Buy Now</a>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<div class="mui-col-xl-3 mui-col-md-3 mui-col-sm-6 mui-col-xs-12">
		<div class="prime-pack rounded">
			<div class="ribbon-B"><span>Prime Pack</span></div>
			<div class="course-card-thumbnail">
				<a href="https://www.tutorialspoint.com/prime-pack/java-prime-pack/index.asp"><img class="rounded-img" src="https://d3mxt5v3yxgcsr.cloudfront.net/courses/4879/course_4879_image.jpg" alt="Java Prime Pack" title="Java Prime Pack">
				<span class="prime-icon-trigger"></span></a>
			</div>
			<div class="primepack-card-body">
				<h4 class="h48"><a href="https://www.tutorialspoint.com/prime-pack/java-prime-pack/index.asp" title="Java Prime Pack">Java Prime Pack</a></h4>
				<p class="videos-details"><i class="fa fa-play-circle"></i> &nbsp;9 Courses &nbsp;&nbsp;
					<span><i class="fa fa-file-pdf"></i> &nbsp;2 eBooks </span>
				</p>
				<p class="h22">Created by - <a href="https://www.tutorialspoint.com/profile/tutorialspoint_com" target="_blank">Tutorialspoint</a></p>
				<div class="ebk-course-price">
					<div class="ebook-price" style="display: block;">
						<div class="">
							<b><span class="price" data-usd="59.99" data-inr="2599.99" data-cad="79.99" data-aud="79.99" data-gbp="59.99" data-eur="59.99" data-jpy="7999.99">₹2,599.99</span></b>
							<span class="actual-price" data-usd="499.99" data-inr="36599.99" data-cad="609.99" data-aud="639.99" data-gbp="359.99" data-eur="414.99" data-jpy="54599.99">₹36,599.99</span>
						</div>						
						<div class="ebook-discount">				
							<span>You Save</span> 
							<b><span class="discount-amount" data-usd="440" data-inr="34000" data-cad="530" data-aud="560" data-gbp="300" data-eur="355" data-jpy="46600">₹34,000.00</span></b>
							<span class="ebook-off rounded" data-usd="88" data-inr="93" data-cad="87" data-aud="88" data-gbp="83" data-eur="86" data-jpy="85" style="display: block;">93 % off</span> 		
						</div>
					</div>
					<div class="mui-col-md-12 mui-col-sm-12 add_to_cart_div np">
						<div class="row">
							<div class="btn-success rounded">
								<a href="https://www.tutorialspoint.com/prime-pack/java-prime-pack/index.asp" target="_blank">Buy Now</a>
							</div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	</div>
	<div class="clear"></div>
	<div class="col-md-12 text-right pr-10">
		<a href="https://www.tutorialspoint.com/latest/prime-packs" title="card view" class="btn btn-viewmore">view more <i class="fa fa-external-link"></i> </a>
	</div>
</div>

<!-- video courses, ebooks and prime pack content starts here -->
<div class="mui-container videos-wrap np">
	<h1 class="uk-heading-line"><span><i class="fa fa-video"></i> Video <b>Courses</b></span></h1>
	<div class="slides">
	<div class="mui-col-xl-3 mui-col-md-3 mui-col-sm-6 mui-col-xs-12">
		<div class="course-card video-card rounded">
			<div class="ribbon-V"><span>Video</span></div>
			<div class="course-card-thumbnail">
				<a href="https://www.tutorialspoint.com/the_full_stack_web_development/index.asp" target="_blank"><img class="rounded-img" src="https://d3mxt5v3yxgcsr.cloudfront.net/courses/2265/course_2265_image.png" alt="Full Stack Web Development" title="Full Stack Web Development">
				<span class="play-button-trigger"></span></a>
			</div>
			<div class="course-card-body">
				<h4 class="h48"><a href="https://www.tutorialspoint.com/the_full_stack_web_development/index.asp" title="The Full Stack Web Development">The Full Stack Web Development</a></h4>
				<p class="videos-details course-list-video">207 Lectures <span><i class="fa fa-clock"></i> 33 hours</span></p>
				<p class="h22"><i class="fa fa-user"></i><a href="https://www.tutorialspoint.com/profile/eduonix_learning_solutions"> Eduonix Learning Solutions</a></p>
				<div class="v-courses-price">
					<div class="ebook-price-card">
						<div class="ebook-price" style="display: block;">
							<div class="">
								<b><span class="price" data-usd="8.99" data-inr="399.99" data-cad="10.99" data-aud="11.49" data-gbp="8.99" data-eur="8.99" data-jpy="977.39">₹399.99</span></b>
								<span class="actual-price" data-usd="184" data-inr="12880" data-cad="224.48" data-aud="235.52" data-gbp="132.48" data-eur="152.72" data-jpy="20004.48">₹12,880.00</span>
							</div>						
							<div class="ebook-discount">				
								<span>You Save</span> 
								<b><span class="discount-amount" data-usd="175.01" data-inr="12480.01" data-cad="213.49" data-aud="224.03" data-gbp="123.49" data-eur="143.73" data-jpy="19027.09">₹12,480.01</span></b>
								<span class="ebook-off rounded" data-usd="95" data-inr="97" data-cad="95" data-aud="95" data-gbp="93" data-eur="94" data-jpy="95" style="display: block;">97 % off</span> 		
							</div>
						</div>
					</div>
					<div class="mui-col-md-12 mui-col-sm-12 add_to_cart_div np">
						<div class="row">
							<div class="btn-success rounded"><a href="https://www.tutorialspoint.com/the_full_stack_web_development/index.asp" target="_blank">Buy Now</a></div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<div class="mui-col-xl-3 mui-col-md-3 mui-col-sm-6 mui-col-xs-12">
		<div class="course-card video-card rounded">
			<div class="ribbon-V"><span>Video</span></div>
			<div class="course-card-thumbnail">
				<a href="https://www.tutorialspoint.com/angular_8_full_stack_development_with_spring_boot/index.asp" target="_blank"><img class="rounded-img" src="https://d3mxt5v3yxgcsr.cloudfront.net/courses/2373/course_2373_image.jpg" alt="Angular 8 Full Stack Development With Spring Boot" title="Angular 8 Full Stack Development With Spring Boot">
				<span class="play-button-trigger"></span></a>
			</div>
			<div class="course-card-body">
				<h4 class="h48"><a href="https://www.tutorialspoint.com/angular_8_full_stack_development_with_spring_boot/index.asp" title="Angular 8 Full Stack Development With Spring Boot">Angular 8 Full Stack Development With Spring Boot</a></h4>
				<p class="videos-details course-list-video">69 Lectures <span><i class="fa fa-clock"></i> 5 hours</span></p>
				<p class="h22"><i class="fa fa-user"></i> <a href="https://www.tutorialspoint.com/profile/senol_atac"> Senol Atac</a></p>
				<div class="v-courses-price">
					<div class="ebook-price-card">
						<div class="ebook-price" style="display: block;">
							<div class="">
								<b><span class="price" data-usd="8.99" data-inr="399.99" data-cad="10.99" data-aud="11.49" data-gbp="8.99" data-eur="8.99" data-jpy="977.39">₹399.99</span></b>
								<span class="actual-price" data-usd="130" data-inr="9360" data-cad="158.6" data-aud="166.4" data-gbp="93.6" data-eur="107.9" data-jpy="14133.6">₹9,360.00</span>
							</div>						
							<div class="ebook-discount">				
								<span>You Save</span> 
								<b><span class="discount-amount" data-usd="121.01" data-inr="8960.01" data-cad="147.61" data-aud="154.91" data-gbp="84.61" data-eur="98.91" data-jpy="13156.21">₹8,960.01</span></b>
								<span class="ebook-off rounded" data-usd="93" data-inr="96" data-cad="93" data-aud="93" data-gbp="90" data-eur="92" data-jpy="93" style="display: block;">96 % off</span> 		
							</div>
						</div>
					</div>
					<div class="mui-col-md-12 mui-col-sm-12 add_to_cart_div np">
						<div class="row">
							<div class="btn-success rounded"><a href="https://www.tutorialspoint.com/angular_8_full_stack_development_with_spring_boot/index.asp" target="_blank">Buy Now</a></div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<div class="mui-col-xl-3 mui-col-md-3 mui-col-sm-6 mui-col-xs-12">
		<div class="course-card video-card rounded">
			<div class="ribbon-V"><span>Video</span></div>
			<div class="course-card-thumbnail">
				<a href="https://www.tutorialspoint.com/python_for_beginners_learn_python_programming_python_3/index.asp" target="_blank"><img class="rounded-img" src="https://d3mxt5v3yxgcsr.cloudfront.net/courses/3211/course_3211_image.jpg" alt="Python For Beginners: Learn Python Programming (Python 3)" title="Python For Beginners: Learn Python Programming (Python 3)">
				<span class="play-button-trigger"></span></a>
			</div>
			<div class="course-card-body">
				<h4 class="h48"><a href="https://www.tutorialspoint.com/python_for_beginners_learn_python_programming_python_3/index.asp" title="Python For Beginners: Learn Python Programming (Python 3)" target="_blank">Python For Beginners: Learn Python Programming (Python 3)</a></h4>
				<p class="videos-details course-list-video">60 Lectures <span><i class="fa fa-clock"></i> 2 hours</span></p>
				<p class="h22"><i class="fa fa-user"></i> <a href="https://www.tutorialspoint.com/profile/jason_cannon"> Jason Cannon</a></p>
				<div class="v-courses-price">
					<div class="ebook-price" style="display: block;">
						<div class="">
							<b><span class="price" data-usd="8.99" data-inr="399.99" data-cad="10.99" data-aud="11.49" data-gbp="8.99" data-eur="8.99" data-jpy="977.39">₹399.99</span></b>
							<span class="actual-price" data-usd="130" data-inr="9100" data-cad="158.6" data-aud="166.4" data-gbp="93.6" data-eur="107.9" data-jpy="14133.6">₹9,100.00</span>
						</div>						
						<div class="ebook-discount">				
							<span>You Save</span> 
							<b><span class="discount-amount" data-usd="121.01" data-inr="8700.01" data-cad="147.61" data-aud="154.91" data-gbp="84.61" data-eur="98.91" data-jpy="13156.21">₹8,700.01</span></b>
							<span class="ebook-off rounded" data-usd="93" data-inr="96" data-cad="93" data-aud="93" data-gbp="90" data-eur="92" data-jpy="93" style="display: block;">96 % off</span> 		
						</div>
					</div>
					<div class="mui-col-md-12 mui-col-sm-12 add_to_cart_div np">
						<div class="row">
							<div class="btn-success rounded"><a href="https://www.tutorialspoint.com/python_for_beginners_learn_python_programming_python_3/index.asp" target="_blank"> Buy Now </a></div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	<div class="mui-col-xl-3 mui-col-md-3 mui-col-sm-6 mui-col-xs-12">
		<div class="course-card video-card rounded">
			<div class="ribbon-V"><span>Video</span></div>
			<div class="course-card-thumbnail">
				<a href="https://www.tutorialspoint.com/machine-learning-interview-questions-and-answers/index.asp" target="_blank"><img class="rounded-img" src="https://d3mxt5v3yxgcsr.cloudfront.net/courses/5978/course_5978_image.jpg" alt="Machine Learning Interview Questions & Answers" title="Machine Learning Interview Questions & Answers" />
				<span class="play-button-trigger"></span></a>
			</div>
			<div class="course-card-body">
				<h4 class="h48"><a href="https://www.tutorialspoint.com/machine-learning-interview-questions-and-answers/index.asp" title="Machine Learning Interview Questions & Answers" target="_blank">Machine Learning Interview Questions & Answers</a></h4>
				<p class="videos-details course-list-video">11 Lectures <span><i class="fa fa-clock"></i> 3 hours</span></p>
				<p class="h22"><i class="fa fa-user"></i> <a href="https://www.tutorialspoint.com/profile/uplatz-163782114078"> Uplatz </a></p>
				<div class="v-courses-price">
					<div class="ebook-price" style="display: block;">
						<div class="">
							<b><span class="price" data-usd="8.99" data-inr="399.99" data-cad="10.99" data-aud="11.49" data-gbp="8.99" data-eur="8.99">₹399.99</span></b>
							<span class="actual-price" data-usd="100" data-inr="7337" data-cad="122" data-aud="128" data-gbp="72" data-eur="83">₹7,337.00</span>
						</div>						
						<div class="ebook-discount">				
							<span>You Save</span> 
							<b><span class="discount-amount" data-usd="91.01" data-inr="6937.01" data-cad="111.01" data-aud="116.51" data-gbp="63.01" data-eur="74.01">₹6,937.01</span></b>
							<span class="ebook-off rounded" data-usd="91" data-inr="95" data-cad="91" data-aud="91" data-gbp="88" data-eur="89" style="display: block;">95 % off</span> 		
						</div>
					</div>
					<div class="mui-col-md-12 mui-col-sm-12 add_to_cart_div np">
						<div class="row">
							<div class="btn-success rounded"><a href="https://www.tutorialspoint.com/machine-learning-interview-questions-and-answers/index.asp" target="_blank"> Buy Now </a></div>
						</div>
					</div>
				</div>
			</div>
		</div>
	</div>
	</div>
	<div class="clear"></div>
	<div class="col-md-12 text-right pr-10">
		<a href="https://www.tutorialspoint.com/market/index.asp?v=1.00" title="card view" class="btn btn-viewmore">view more <i class="fa fa-external-link"></i> </a>
	</div>
</div>
<!-- end of video courses, ebooks and prime pack content starts here -->

<div class="clear"></div>
<div class="mui-container tp-home">
	<h1><a href="https://www.tutorialspoint.com/tutorialslibrary.htm" title="tutorialspoint library"><i class="fa fa-cubes"></i> Tutorials <b>Library</b> <br/></a></h1>   
	
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-cloud-upload-alt"></i> Latest Technologies</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 bitcoin library-np">
					<a href="https://www.tutorialspoint.com/bitcoin/index.htm" target="_blank" title="Bitcoin">
						<div class="imgs-title">Bitcoin</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 blockchain library-np">
					<a href="https://www.tutorialspoint.com/blockchain/index.htm" target="_blank" title="Blockchain">
						<div class="imgs-title">Blockchain</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 blueprism library-np">
					<a href="https://www.tutorialspoint.com/blue_prism/index.htm" target="_blank" title="Blue Prism Tutorial">
						<div class="imgs-title">Blue Prism</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 ethereum library-np">
					<a href="https://www.tutorialspoint.com/ethereum/index.htm" target="_blank" title="Ethereum Tutorial">
						<div class="imgs-title">Ethereum</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 openshift library-np">
					<a href="https://www.tutorialspoint.com/openshift/index.htm" target="_blank" title="OpenShift Tutorial">
						<div class="imgs-title">OpenShift</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 python library-np">
					<a href="https://www.tutorialspoint.com/python_blockchain/index.htm" target="_blank" title="Python Blockchain Tutorial">
						<div class="imgs-title">Python Blockchain</div>
					</a>
				</div>
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/latest_technologies.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-check-square"></i> Machine Learning</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 machine-learn library-np">
					<a href="https://www.tutorialspoint.com/machine_learning/index.htm" target="_blank" title="Machine Learning Tutorial">
						<div class="imgs-title">Machine Learning</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 tensorflow library-np">
					<a href="https://www.tutorialspoint.com/tensorflow/index.htm" target="_blank" title="TensorFlow Tutorial">
						<div class="imgs-title">TensorFlow</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 machine-learn-python library-np">
					<a href="https://www.tutorialspoint.com/machine_learning_with_python/index.htm" target="_blank" title="Machine Learning with Python Tutorial">
						<div class="imgs-title">ML with Python</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 ai-python library-np">
					<a href="https://www.tutorialspoint.com/artificial_intelligence_with_python/index.htm" target="_blank" title="AI with Python Tutorial">
						<div class="imgs-title">AI with Python</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 timeseries library-np">
					<a href="https://www.tutorialspoint.com/time_series/index.htm" target="_blank" title="Time Series Tutorial">
						<div class="imgs-title">Time Series</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 pytorch library-np">
					<a href="https://www.tutorialspoint.com/pytorch/index.htm" target="_blank" title="PyTorch Tutorial">
						<div class="imgs-title">PyTorch</div>
					</a>
				</div>
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/machine_learning_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-cogs"></i> Computer Science</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 computer-fundamentals library-np">
					<a href="https://www.tutorialspoint.com/computer_fundamentals/index.htm" target="_blank" title="Computer-Fundamentals Tutorial">
						<div class="imgs-title">Computer Fundamentals</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 complier-design library-np">
					<a href="https://www.tutorialspoint.com/compiler_design/index.htm" target="_blank" title="Compiler-Design Tutorial">
						<div class="imgs-title">Compiler Design</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 operating-system library-np">
					<a href="https://www.tutorialspoint.com/operating_system/index.htm" target="_blank" title="Operating-System Tutorial">
						<div class="imgs-title">Operating System</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 data-structures library-np">
					<a href="https://www.tutorialspoint.com/data_structures_algorithms/index.htm" target="_blank" title="Data Structure Tutorial">
						<div class="imgs-title">Data Structure</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 dbms library-np">
					<a href="https://www.tutorialspoint.com/dbms/index.htm" target="_blank" title="DBMS Tutorial">
						<div class="imgs-title">DBMS</div> 
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 networking library-np">
					<a href="https://www.tutorialspoint.com/data_communication_computer_network/index.htm" target="_blank" title="Networking Tutorial">
						<div class="imgs-title">Networking</div>
					</a>
				</div>
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/computer_science_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-globe"></i> Web Development</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 html library-np">
					<a href="https://www.tutorialspoint.com/html/index.htm" target="_blank" title="html Tutorial">
						<div class="imgs-title">HTML</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 css library-np">
					<a href="https://www.tutorialspoint.com/css/index.htm" target="_blank" title="css Tutorial">
						<div class="imgs-title">css</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 javascript library-np">
					<a href="https://www.tutorialspoint.com/javascript/index.htm" target="_blank" title="Javascript Tutorial">
						<div class="imgs-title">Javascript</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 php library-np">
					<a href="https://www.tutorialspoint.com/php/index.htm" target="_blank" title="PHP Tutorial">
						<div class="imgs-title">PHP</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 angularjs4 library-np">
					<a href="https://www.tutorialspoint.com/angular4/index.htm" target="_blank" title="Angular JS 4 Tutorial">
						<div class="imgs-title">Angular JS 4</div> 
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 mysql library-np">
					<a href="https://www.tutorialspoint.com/mysql/index.htm" target="_blank" title="MySQL Tutorial">
						<div class="imgs-title">MySQL</div>
					</a>
				</div>
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/web_development_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-code"></i> Programming Tutorials</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 c-programming library-np">
					<a href="https://www.tutorialspoint.com/cprogramming/index.htm" target="_blank" title="C Programming Tutorial">
						<div class="imgs-title">C Programming</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 cplus library-np">
					<a href="https://www.tutorialspoint.com/cplusplus/index.htm" target="_blank" title="C++  Tutorial">
						<div class="imgs-title">C++ </div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 java-8 library-np">
					<a href="https://www.tutorialspoint.com/java8/index.htm" target="_blank" title="Java 8 Tutorial">
						<div class="imgs-title">Java 8</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 python-n library-np">
					<a href="https://www.tutorialspoint.com/python/index.htm" target="_blank" title="Python Tutorial">
						<div class="imgs-title">Python</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 scala library-np">
					<a href="https://www.tutorialspoint.com/scala/index.htm" target="_blank" title="Scala Tutorial">
						<div class="imgs-title">Scala</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 chash library-np">
					<a href="https://www.tutorialspoint.com/csharp/index.htm" target="_blank" title="C# Tutorial">
						<div class="imgs-title">C#</div> 
					</a>
				</div>
				
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/computer_programming_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-java"></i> Java Technologies</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 java-8 library-np">
					<a href="https://www.tutorialspoint.com/java8/index.htm" target="_blank" title="Java 8 Tutorial">
						<div class="imgs-title">Java 8</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 jdbc library-np">
					<a href="https://www.tutorialspoint.com/jdbc/index.htm" target="_blank" title="JDBC  Tutorial">
						<div class="imgs-title">JDBC </div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 swing library-np">
					<a href="https://www.tutorialspoint.com/swing/index.htm" target="_blank" title="Swing Tutorial">
						<div class="imgs-title">Swing</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 servlets library-np">
					<a href="https://www.tutorialspoint.com/servlets/index.htm" target="_blank" title="Servlets Tutorial">
						<div class="imgs-title">Servlets</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 spring library-np">
					<a href="https://www.tutorialspoint.com/spring/index.htm" target="_blank" title="Spring Tutorial">
						<div class="imgs-title">Spring</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 hibernate library-np">
					<a href="https://www.tutorialspoint.com/hibernate/index.htm" target="_blank" title="Hibernate Tutorial">
						<div class="imgs-title">Hibernate</div> 
					</a>
				</div>				
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/java_technology_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-mobile"></i> Mobile App Development</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 android library-np">
					<a href="https://www.tutorialspoint.com/android/index.htm" target="_blank" title="Android Tutorial">
						<div class="imgs-title">Android</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 swift library-np">
					<a href="https://www.tutorialspoint.com/swift/index.htm" target="_blank" title="Swift Tutorial">
						<div class="imgs-title">Swift </div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 ios library-np">					
					<a href="https://www.tutorialspoint.com/ios/index.htm" target="_blank" title="IOS Tutorial">
						<div class="imgs-title">IOS</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 kotlin library-np">
					<a href="https://www.tutorialspoint.com/kotlin/index.htm" target="_blank" title="Kotlin Tutorial">
						<div class="imgs-title">Kotlin</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 react library-np">
					<a href="https://www.tutorialspoint.com/react_native/index.htm" target="_blank" title="React Native Tutorial">
						<div class="imgs-title">React Native</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 xamarian library-np">
					<a href="https://www.tutorialspoint.com/xamarin/index.htm" target="_blank" title="Xamarian Tutorial">
						<div class="imgs-title">Xamarian</div> 
					</a>
				</div>				
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/mobile_development_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-database"></i> Database Tutorials</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 mongo-db library-np">
					<a href="https://www.tutorialspoint.com/mongodb/index.htm" target="_blank" title="Mongo DB Tutorial">
						<div class="imgs-title">Mongo DB</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 pl-sql library-np">
					<a href="https://www.tutorialspoint.com/plsql/index.htm" target="_blank" title="PL / SQL Tutorial">
						<div class="imgs-title">PL / SQL </div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 sql library-np">					
					<a href="https://www.tutorialspoint.com/sql/index.htm" target="_blank" title="SQL Tutorial">
						<div class="imgs-title">SQL</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 db2 library-np">
					<a href="https://www.tutorialspoint.com/db2/index.htm" target="_blank" title="DB2 Tutorial">
						<div class="imgs-title">DB2</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 mysql library-np">
					<a href="https://www.tutorialspoint.com/mysql/index.htm" target="_blank" title="MySQL Tutorial">
						<div class="imgs-title">MySQL</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 memcached library-np">
					<a href="https://www.tutorialspoint.com/memcached/index.htm" target="_blank" title="Memcached Tutorial">
						<div class="imgs-title">Memcached</div> 
					</a>
				</div>				
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/database_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-windows"></i> Microsoft Technologies</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 asp-net library-np">
					<a href="https://www.tutorialspoint.com/asp.net/index.htm" target="_blank" title="ASP.NET Tutorial">
						<div class="imgs-title">ASP.NET</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 entity-framework library-np">
					<a href="https://www.tutorialspoint.com/entity_framework/index.htm" target="_blank" title="Entity Framework Tutorial">
						<div class="imgs-title">Entity Framework</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 vb-net library-np">					
					<a href="https://www.tutorialspoint.com/vb.net/index.htm" target="_blank" title="VB.NET Tutorial">
						<div class="imgs-title">VB.NET</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 microsoft-project library-np">
					<a href="https://www.tutorialspoint.com/ms_project/index.htm" target="_blank" title="Microsoft Project Tutorial">
						<div class="imgs-title">Microsoft Project</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 microsoft-excel library-np">
					<a href="https://www.tutorialspoint.com/excel/index.htm" target="_blank" title="Microsoft Excel Tutorial">
						<div class="imgs-title">Microsoft Excel</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 microsoft-word library-np">
					<a href="https://www.tutorialspoint.com/word/index.htm" target="_blank" title="Microsoft Word Tutorial">
						<div class="imgs-title">Microsoft Word</div> 
					</a>
				</div>				
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/microsoft_technologies_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	<div class="mui-col-md-6 mui-col-sm-12 mui-col-xs-12">
		<div class="tp-content-widget">
			<h2><i class="fa fa-database"></i> Big Data and Analytics</h2>
			<div class="library-sprite">
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 big-data-analytics library-np">
					<a href="https://www.tutorialspoint.com/big_data_analytics/index.htm" target="_blank" title="Big Data Analytics Tutorial">
						<div class="imgs-title">Big Data Analytics</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 hadoop library-np">
					<a href="https://www.tutorialspoint.com/hadoop/index.htm" target="_blank" title="Hadoop Tutorial">
						<div class="imgs-title">Hadoop</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 sas library-np">					
					<a href="https://www.tutorialspoint.com/sas/index.htm" target="_blank" title="SAS Tutorial">
						<div class="imgs-title">SAS</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 qlikview library-np">
					<a href="https://www.tutorialspoint.com/qlikview/index.htm" target="_blank" title="QlikView Tutorial">
						<div class="imgs-title">QlikView</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 powerbi library-np">
					<a href="https://www.tutorialspoint.com/power_bi/index.htm" target="_blank" title="Power BI Tutorial">
						<div class="imgs-title">Power BI</div>
					</a>
				</div>
				<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6 tableau library-np">
					<a href="https://www.tutorialspoint.com/tableau/index.htm" target="_blank" title="Tableau Tutorial">
						<div class="imgs-title">Tableau</div> 
					</a>
				</div>				
				<div class="clear"></div>
				<div class="tp-content-viewmore"><a href="https://www.tutorialspoint.com/big_data_tutorials.htm"><i class="fa fa-external-link"></i> view all</a></div>
			</div>
		</div>
	</div>
	
	<div class="mui-col-md-12 mui-col-sm-12 mui-col-xs-12 tp-browse-more">
		<a href="https://www.tutorialspoint.com/tutorialslibrary.htm" target="_blank" title="Browse Complete Library"> &nbsp; Browse Complete Library <i class="fa fa-books"></i></a> 
	</div>
</div>
<div class="clear"></div>
		
<div class="tp-coding-ground">   
	<div class="mui-container tp-home" style="padding-left:0px; padding-right:0px;">
		<h1><a href="https://www.tutorialspoint.com/codingground.htm" target="_blank" title="Coding Ground - Free Online IDE and Terminal"><i class="fa fa-code"></i> Coding <b>Ground</b></a></h1>
		<p>Coding Platform For Your Website <br/>
		Available for <b>75+</b> Programming Languages</p>
		<a href="https://www.tutorialspoint.com/coding_platform_for_websites.htm" target="_blank" title="Coding Platform for Websites" class="tp-coding-btn">How it works?</a>
	</div>
</div>
	
<div class="mui-container tp-home">
	<h1><i class="fa fa-book-open"></i> Selected <b>Reading</b></h1>
	<br/>
	<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6">
		<div class="select-reading-box">
			<a href="https://www.tutorialspoint.com/developers_best_practices/index.htm" title="Developer's Best Practices">
				<img src="https://www.tutorialspoint.com/images/developers-best-practices.png" alt="Developer's Best Practices" title="Developer's Best Practices">
			<div class="">Developers <br/><b>Best Practices</b></div></a>
		</div>
	</div>
	<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6">
		<div class="select-reading-box">
			<a href="https://www.tutorialspoint.com/effective_resume_writing.htm" title="Effective Resume Writing">
				
			<div class="">Effective <br/><b>Resume Writing</b></div></a>
		</div>
	</div>
	<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6">
		<div class="select-reading-box">
			<a href="https://www.tutorialspoint.com/computer_glossary.htm" title="Computer Glossary">
				<img src="https://www.tutorialspoint.com/images/computer-glossary.png" alt="Computer Glossary" title="Computer Glossary">
			<div class="">Computer <br/><b>Glossary</b></div></a>
		</div>
	</div>
	<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6">
		<div class="select-reading-box">
			<a href="https://www.tutorialspoint.com/computer_whoiswho.htm" title="Who is Who in Computer">
				
			<div class=""><b>Who is Who</b> <br/>in Computer</div></a>
		</div>
	</div>
	<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6">
		<div class="select-reading-box">
			<a href="https://www.tutorialspoint.com/questions_and_answers.htm" title="Technical Questions and Answers">
				<img src="https://www.tutorialspoint.com/images/questions-answers.png" alt="Technical Questions and Answers" title="Technical Questions and Answers">
			<div class="">Technical <br/><b>Q & A</b></div></a>
		</div>
	</div>
	<div class="mui-col-md-2 mui-col-sm-6 mui-col-xs-6">
		<div class="select-reading-box">
			<a href="https://www.tutorialspoint.com/multi_language_tutorials.htm" title="Multi-Lingual Tutorials">
				<img src="https://www.tutorialspoint.com/images/multilanguage-tutorials.png" alt="Multi-Lingual Tutorials" title="Multi-Lingual Tutorials">
			<div class=""><b>Multi-Lingual</b> <br/>Tutorials</div></a>
		</div>
	</div>
</div>   
<!-- End of Body Content -->
	
<div class="footer-nav">
	<div class="mui-container">
		<div class="mui-col-md-3 mui-col-sm-6 mui-col-xs-6">
			<h3>Mobile First</h3>
			<ul class="">
				<li><a class="iphone" href="https://itunes.apple.com/us/app/tutorials-point/id914891263?ls=1&amp;mt=8" target="_blank" rel="nofollow"><img alt="tutorialspoint iPhone App" src="https://www.tutorialspoint.com/images/apple_store.jpg" title="tutorialspoint iPhone App" class="mobile-apps"></a></li>
				<li><a class="android" href="https://play.google.com/store/apps/details?id=com.tutorialspoint.onlineviewer" target="_blank" rel="nofollow"><img src="https://www.tutorialspoint.com/images/google_play.jpg" alt="tutorialspoint Android app" title="tutorialspoint Android app" class="mobile-apps"></a></li>
				<li><a class="microsoft" href="http://www.windowsphone.com/s?appid=91249671-7184-4ad6-8a5f-d11847946b09" target="_blank" rel="nofollow"><img src="https://www.tutorialspoint.com/images/windows_store.jpg" alt="tutorialspoint Windows app" title="tutorialspoint Windows app" class="mobile-apps"></a></li>
			</ul>
		</div>
		<div class="mui-col-md-3 mui-col-sm-6 mui-col-xs-6">
			<h3>About Us</h3>
			<ul class="">
				<li><i class="fa fa-check"></i> <a href="/about/index.htm">Company</a></li>
				<li><i class="fa fa-check"></i> <a href="/about/about_team.htm">Our Team</a></li>
				<li><i class="fa fa-check"></i> <a href="/about/about_careers.htm">Careers</a></li>
				<li><i class="fa fa-check"></i> <a href="/about/about_privacy.htm">Privacy Policy</a></li>
				<li><i class="fa fa-check"></i> <a href="/about/about_cookies.htm">Cookies Policy</a></li>
				<li><i class="fa fa-check"></i> <a href="/about/about_terms_of_use.htm">Terms of use</a></li>
				<li><i class="fa fa-check"></i> <a href="/about/return_refund_policy.htm">Refund Policy</a></li>
			</ul>
		</div>
		<div class="mui-col-md-3 mui-col-sm-6 mui-col-xs-6 foot-clear">
			<h3>Extra Links</h3>
			<ul class="">
				<li><i class="fa fa-check"></i> <a href="https://www.tutorialspoint.com/online_dev_tools.htm">Dev Tools</a></li>
				<li><i class="fa fa-check"></i> <a href="https://www.tutorialspoint.com/free_web_graphics.htm">Free Graphics</a></li>
				<li><i class="fa fa-check"></i> <a href="https://www.tutorialspoint.com/online_file_conversion.htm">File Conversion</a></li>
				<li><i class="fa fa-check"></i> <a href="https://www.tutorialspoint.com/netmeeting.php">NetMeeting</a></li>
				<li><i class="fa fa-check"></i> <a href="https://www.tutorialspoint.com/free_online_whiteboard.htm">Whiteboard</a></li>
			</ul>
		</div>
		<div class="mui-col-md-3 mui-col-sm-6 mui-col-xs-6">
			<h3>Contact Us</h3>
			<ul>
				<li><p><i class="icon icon-map-marker"></i> <strong>Address:</strong> 4th Floor, Incor9 Building, Kavuri Hills, Madhapur, Hyderabad, Telangana 500081</p></li>
				<li><p><i class="icon icon-dribbble"></i> <strong>Website:</strong> <a href="http://www.tutorialspoint.com">www.tutorialspoint.com</a></p></li>
			</ul>
		</div>
	</div>
</div>
	
<div class="footer-copyright">
	<div class="mui-container tp-footer-social-links">
		<div class="">
			<div class="mui-col-md-4">
				<h3>Follow</h3>
				<a title="Follow us on Facebook" rel="nofollow" href="https://www.facebook.com/tutorialspointindia" class="facebook" target="_blank"><i class="fa fa-facebook-f"></i> </a>
				<a title="Follow us on Google+" rel="nofollow" href="https://plus.google.com/u/0/+tutorialspoint" class="google" target="_blank"><i class="fa fa-google-plus-g"></i> </a>
				<a title="Follow us on Twitter" rel="nofollow" href="http://www.twitter.com/tutorialspoint" class="twitter" target="_blank"><i class="fa fa-twitter"></i> </a>
				<a title="Follow us on Linkedin" rel="nofollow" href="http://www.linkedin.com/company/tutorialspoint" class="linkedin" target="_blank"><i class="fa fa-linkedin-in"></i> </a>         
				<a title="Follow us on Youtube" rel="nofollow" href="https://www.youtube.com/channel/UCVLbzhxVTiTLiVKeGV7WEBg" class="linkedin" target="_blank"><i class="fa fa-youtube"></i> </a>         
			</div>
			<div class="mui-col-md-4">
				<a href="https://www.tutorialspoint.com/index.htm" class="logo"></a>
				<p><span>&copy; Copyright 2021. All Rights Reserved.</span></p>
			</div>
			<div class="mui-col-md-4 news-group">
				<h3 style="padding:0px 0px 0px 0px;">Newsletter</h3>
				<div class="">
					<input type="text" class="form-control-foot search" name="textemail" id="textemail" autocomplete="off" placeholder="Enter email for newsletter" onfocus="if (this.value == 'Enter email for newsletter...') { this.value = ''; }" onblur="if (this.value == '') {this.value = 'Enter email for newsletter...'; }">
					<span class="input_group_button"> <button class="btn btn-default btn-footer" id="btnemail" type="submit" onclick="javascript:void(0);">go</button> </span>
					<div id="newsresponse" style="display:none;"></div>
				</div>         
				<ul>
					<li><a href="/about/about_privacy.htm#cookies"><i class="fa fa-globe"></i> Cookies Policy</a></li>
					<li><a href="/about/faq.htm"><i class="fa fa-question-circle"></i> FAQ's</a></li>
					<li><a href="/about/about_helping.htm"><i class="fa fa-hands-helping"></i> Helping</a></li>
					<li><a href="/about/contact_us.htm"><i class="fa fa-user"></i> Contact</a></li>
				</ul>         
			</div>
		</div>
	</div>
</div>
<div id="privacy-banner">
  <div>
    <p>
      We make use of cookies to improve our user experience. By using this website, you agree with our Cookies Policy.
      <a id="banner-accept" href="javascript:void(0)">Agree</a>
      <a id="banner-learn" href="/about/about_cookies.htm" target="_blank">Learn more</a>
    </p>
  </div>
</div>
<!-- Global site tag (gtag.js) - Google Analytics -->
<script src="/themes/home/mui-combined.min.js"></script>
<script src="/themes/home/jquery.min.js"></script>
<script src="/themes/home/fontawesome-home-min.js?v=1.223"></script>
<script src="https://www.tutorialspoint.com/theme/js/localization.js?v=2.2"></script>
<script src="/themes/js/home.js?v=5"></script>
<script async src="https://www.googletagmanager.com/gtag/js?id=UA-232293-6"></script>
<script>
	window.dataLayer = window.dataLayer || [];
	function gtag(){dataLayer.push(arguments);}
	gtag('js', new Date());
	gtag('config', 'UA-232293-6');
</script>
</body>
</html>
`
	result := make([]string, 0)
	var wg sync.WaitGroup

	wg.Add(10)
	doc, err := html.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, img := range n.Attr {
				if img.Key == "src" {
					result = append(result, img.Val)

				}
			}

		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	//fmt.Println(result)

	for _, result := range result {
		go func(result string) {
			defer wg.Done()
			tokens := strings.Split(result, "/")
			imageName := tokens[len(tokens)-1]
			fmt.Println("Downloading", result, "to", imageName)

			output, err := os.Create(imageName)
			if err != nil {
				log.Fatal(err)
			}
			defer output.Close()

			res, err := http.Get(result)
			if err != nil {
				log.Fatal(err)
			} else {
				defer res.Body.Close()
				_, err = io.Copy(output, res.Body)
				if err != nil {
					log.Fatal(err)
				} else {
					fmt.Println("Downloaded", imageName)
				}
			}
		}(result)
	}
	wg.Wait()
	fmt.Println("Done")

}