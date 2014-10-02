## Forwader

The application allows simple forwarding of shorter URLs to specified longer URL's. A url can be set by going to `www.example.com/set/` which will be accessable at `www.example.com`. Multiple URL's can be set at once. For example setting a URL at `www.example.com/set/5` will be accessable via `www.exmaple.com/5`. This currently only works via numbers.

## Requirements

* Server with PHP >= 5.4 and MCrypt PHP Extension
* [Composer](https://getcomposer.org/)
* [Node.JS](http://nodejs.org/)
* [Gulp](http://gulpjs.com/)

## Installation

To install git clone the repo into a directory. Run `composer update` to install PHP dependancies. The SCSS also needs to be compiled. Run `npm install` then `gulp` to compile the SCSS into CSS.

The database details also need to be entered. Rename the file `app/config/database-default.php` file to `database.php` and change the options to your database set-up.

## External Projects Used

* [Laravel](http://laravel.com/) - PHP Framework
* [Jeet](http://jeet.gs/) - CSS Grid System
