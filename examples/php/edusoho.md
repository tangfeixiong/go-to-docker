

[vagrant@kubedev-172-17-4-59 edusoho]$ php /Users/fanhongling/Downloads/99-mirror/php/composer.phar install
Loading composer repositories with package information
Installing dependencies (including require-dev) from lock file
Your requirements could not be resolved to an installable set of packages.

  Problem 1
    - endroid/qrcode 1.8.0 requires ext-gd * -> the requested PHP extension gd is missing from your system.
    - endroid/qrcode 1.8.0 requires ext-gd * -> the requested PHP extension gd is missing from your system.
    - Installation request for endroid/qrcode 1.8.0 -> satisfiable by endroid/qrcode[1.8.0].

  To enable extensions, verify that they are enabled in your .ini files:
    - /etc/php.ini
    - /etc/php.d/20-bz2.ini
    - /etc/php.d/20-calendar.ini
    - /etc/php.d/20-ctype.ini
    - /etc/php.d/20-curl.ini
    - /etc/php.d/20-dom.ini
    - /etc/php.d/20-exif.ini
    - /etc/php.d/20-fileinfo.ini
    - /etc/php.d/20-ftp.ini
    - /etc/php.d/20-gettext.ini
    - /etc/php.d/20-iconv.ini
    - /etc/php.d/20-json.ini
    - /etc/php.d/20-mbstring.ini
    - /etc/php.d/20-pdo.ini
    - /etc/php.d/20-phar.ini
    - /etc/php.d/20-simplexml.ini
    - /etc/php.d/20-sockets.ini
    - /etc/php.d/20-sqlite3.ini
    - /etc/php.d/20-tokenizer.ini
    - /etc/php.d/20-xml.ini
    - /etc/php.d/20-xmlwriter.ini
    - /etc/php.d/20-xsl.ini
    - /etc/php.d/30-pdo_sqlite.ini
    - /etc/php.d/30-wddx.ini
    - /etc/php.d/30-xmlreader.ini
  You can also run `php --ini` inside terminal to see which files are used by PHP in CLI mode.


[vagrant@kubedev-172-17-4-59 edusoho]$ sudo dnf install php-gd                      
Failed to set locale, defaulting to C
Last metadata expiration check: 0:22:44 ago on Tue Mar 20 20:33:48 2018.
Dependencies resolved.
======================================================================================================================================================
 Package                                       Arch                         Version                               Repository                     Size
======================================================================================================================================================
Installing:
 php-gd                                        x86_64                       7.1.15-1.fc26                         updates                        84 k
Installing dependencies:
 aajohan-comfortaa-fonts                       noarch                       3.001-1.fc26                          updates                       146 k
 fontconfig                                    x86_64                       2.12.6-4.fc26                         updates                       249 k
 fontpackages-filesystem                       noarch                       1.44-18.fc26                          fedora                         14 k
 gd                                            x86_64                       2.2.5-1.fc26                          updates                       139 k
 jbigkit-libs                                  x86_64                       2.1-6.fc26                            fedora                         51 k
 libX11                                        x86_64                       1.6.5-2.fc26                          fedora                        614 k
 libX11-common                                 noarch                       1.6.5-2.fc26                          fedora                        165 k
 libXau                                        x86_64                       1.0.8-7.fc26                          fedora                         33 k
 libXpm                                        x86_64                       3.5.12-2.fc26                         fedora                         55 k
 libjpeg-turbo                                 x86_64                       1.5.3-1.fc26                          updates                       152 k
 libtiff                                       x86_64                       4.0.9-1.fc26                          updates                       182 k
 libwebp                                       x86_64                       0.6.1-8.fc26                          updates                       263 k
 libxcb                                        x86_64                       1.12-3.fc26                           fedora                        212 k

Transaction Summary
======================================================================================================================================================
Install  14 Packages

Total download size: 2.3 M
Installed size: 7.0 M
Is this ok [y/N]: y
Downloading Packages:
(1/14): libXpm-3.5.12-2.fc26.x86_64.rpm                                                                               252 kB/s |  55 kB     00:00    
(2/14): php-gd-7.1.15-1.fc26.x86_64.rpm                                                                               277 kB/s |  84 kB     00:00    
(3/14): libX11-common-1.6.5-2.fc26.noarch.rpm                                                                         1.2 MB/s | 165 kB     00:00    
(4/14): libXau-1.0.8-7.fc26.x86_64.rpm                                                                                730 kB/s |  33 kB     00:00    
(5/14): gd-2.2.5-1.fc26.x86_64.rpm                                                                                    1.1 MB/s | 139 kB     00:00    
(6/14): libxcb-1.12-3.fc26.x86_64.rpm                                                                                 817 kB/s | 212 kB     00:00    
(7/14): libjpeg-turbo-1.5.3-1.fc26.x86_64.rpm                                                                         1.6 MB/s | 152 kB     00:00    
(8/14): fontpackages-filesystem-1.44-18.fc26.noarch.rpm                                                               310 kB/s |  14 kB     00:00    
(9/14): libtiff-4.0.9-1.fc26.x86_64.rpm                                                                               1.7 MB/s | 182 kB     00:00    
(10/14): libX11-1.6.5-2.fc26.x86_64.rpm                                                                               759 kB/s | 614 kB     00:00    
(11/14): jbigkit-libs-2.1-6.fc26.x86_64.rpm                                                                           1.0 MB/s |  51 kB     00:00    
(12/14): libwebp-0.6.1-8.fc26.x86_64.rpm                                                                              1.8 MB/s | 263 kB     00:00    
(13/14): fontconfig-2.12.6-4.fc26.x86_64.rpm                                                                          544 kB/s | 249 kB     00:00    
(14/14): aajohan-comfortaa-fonts-3.001-1.fc26.noarch.rpm                                                              441 kB/s | 146 kB     00:00    
------------------------------------------------------------------------------------------------------------------------------------------------------
Total                                                                                                                 678 kB/s | 2.3 MB     00:03     
Running transaction check
Transaction check succeeded.
Running transaction test
Transaction test succeeded.
Running transaction
  Preparing        :                                                                                                                              1/1 
  Installing       : libjpeg-turbo-1.5.3-1.fc26.x86_64                                                                                           1/14 
  Running scriptlet: libjpeg-turbo-1.5.3-1.fc26.x86_64                                                                                           1/14 
  Installing       : fontpackages-filesystem-1.44-18.fc26.noarch                                                                                 2/14 
  Installing       : aajohan-comfortaa-fonts-3.001-1.fc26.noarch                                                                                 3/14 
  Installing       : fontconfig-2.12.6-4.fc26.x86_64                                                                                             4/14 
  Running scriptlet: fontconfig-2.12.6-4.fc26.x86_64                                                                                             4/14 
  Installing       : libwebp-0.6.1-8.fc26.x86_64                                                                                                 5/14 
  Running scriptlet: libwebp-0.6.1-8.fc26.x86_64                                                                                                 5/14 
  Installing       : jbigkit-libs-2.1-6.fc26.x86_64                                                                                              6/14 
  Running scriptlet: jbigkit-libs-2.1-6.fc26.x86_64                                                                                              6/14 
  Installing       : libtiff-4.0.9-1.fc26.x86_64                                                                                                 7/14 
  Running scriptlet: libtiff-4.0.9-1.fc26.x86_64                                                                                                 7/14 
  Installing       : libXau-1.0.8-7.fc26.x86_64                                                                                                  8/14 
  Running scriptlet: libXau-1.0.8-7.fc26.x86_64                                                                                                  8/14 
  Installing       : libxcb-1.12-3.fc26.x86_64                                                                                                   9/14 
  Running scriptlet: libxcb-1.12-3.fc26.x86_64                                                                                                   9/14 
  Installing       : libX11-common-1.6.5-2.fc26.noarch                                                                                          10/14 
  Installing       : libX11-1.6.5-2.fc26.x86_64                                                                                                 11/14 
  Running scriptlet: libX11-1.6.5-2.fc26.x86_64                                                                                                 11/14 
  Installing       : libXpm-3.5.12-2.fc26.x86_64                                                                                                12/14 
  Running scriptlet: libXpm-3.5.12-2.fc26.x86_64                                                                                                12/14 
  Installing       : gd-2.2.5-1.fc26.x86_64                                                                                                     13/14 
  Running scriptlet: gd-2.2.5-1.fc26.x86_64                                                                                                     13/14 
  Installing       : php-gd-7.1.15-1.fc26.x86_64                                                                                                14/14 
  Running scriptlet: php-gd-7.1.15-1.fc26.x86_64                                                                                                14/14 
  Running scriptlet: fontconfig-2.12.6-4.fc26.x86_64                                                                                            14/14 
  Verifying        : php-gd-7.1.15-1.fc26.x86_64                                                                                                 1/14 
  Verifying        : libX11-1.6.5-2.fc26.x86_64                                                                                                  2/14 
  Verifying        : libXpm-3.5.12-2.fc26.x86_64                                                                                                 3/14 
  Verifying        : libX11-common-1.6.5-2.fc26.noarch                                                                                           4/14 
  Verifying        : libxcb-1.12-3.fc26.x86_64                                                                                                   5/14 
  Verifying        : libXau-1.0.8-7.fc26.x86_64                                                                                                  6/14 
  Verifying        : gd-2.2.5-1.fc26.x86_64                                                                                                      7/14 
  Verifying        : libjpeg-turbo-1.5.3-1.fc26.x86_64                                                                                           8/14 
  Verifying        : fontconfig-2.12.6-4.fc26.x86_64                                                                                             9/14 
  Verifying        : fontpackages-filesystem-1.44-18.fc26.noarch                                                                                10/14 
  Verifying        : libtiff-4.0.9-1.fc26.x86_64                                                                                                11/14 
  Verifying        : jbigkit-libs-2.1-6.fc26.x86_64                                                                                             12/14 
  Verifying        : libwebp-0.6.1-8.fc26.x86_64                                                                                                13/14 
  Verifying        : aajohan-comfortaa-fonts-3.001-1.fc26.noarch                                                                                14/14 

Installed:
  php-gd.x86_64 7.1.15-1.fc26                          aajohan-comfortaa-fonts.noarch 3.001-1.fc26          fontconfig.x86_64 2.12.6-4.fc26         
  fontpackages-filesystem.noarch 1.44-18.fc26          gd.x86_64 2.2.5-1.fc26                               jbigkit-libs.x86_64 2.1-6.fc26          
  libX11.x86_64 1.6.5-2.fc26                           libX11-common.noarch 1.6.5-2.fc26                    libXau.x86_64 1.0.8-7.fc26              
  libXpm.x86_64 3.5.12-2.fc26                          libjpeg-turbo.x86_64 1.5.3-1.fc26                    libtiff.x86_64 4.0.9-1.fc26             
  libwebp.x86_64 0.6.1-8.fc26                          libxcb.x86_64 1.12-3.fc26                           

Complete!


[vagrant@kubedev-172-17-4-59 edusoho]$ php /Users/fanhongling/Downloads/99-mirror/php/composer.phar install
Loading composer repositories with package information
Installing dependencies (including require-dev) from lock file
Nothing to install or update
Package endroid/qrcode is abandoned, you should avoid using it. Use endroid/qr-code instead.
Package guzzle/guzzle is abandoned, you should avoid using it. Use guzzlehttp/guzzle instead.
Package phpoffice/phpexcel is abandoned, you should avoid using it. Use phpoffice/phpspreadsheet instead.
Generating autoload files
> Incenteev\ParameterHandler\ScriptHandler::buildParameters
Creating the "app/config/parameters.yml" file
Some parameters are missing. Please provide them.
database_driver (pdo_mysql): 
database_host (127.0.0.1): 
database_port (3306): 
database_name (edusoho): 
database_user (root): 
database_password (null): 
locale (zh_CN): 
secret (ThisTokenIsNotSoSecretChangeIt): 
> Sensio\Bundle\DistributionBundle\Composer\ScriptHandler::buildBootstrap
> Sensio\Bundle\DistributionBundle\Composer\ScriptHandler::clearCache

                                                         
  [Doctrine\DBAL\Exception\DriverException]              
  An exception occured in driver: could not find driver  
                                                         

                                       
  [Doctrine\DBAL\Driver\PDOException]  
  could not find driver                
                                       

                         
  [PDOException]         
  could not find driver  
                         

Script Sensio\Bundle\DistributionBundle\Composer\ScriptHandler::clearCache handling the post-install-cmd event terminated with an exception

                                                                             
  [RuntimeException]                                                         
  An error occurred when executing the "'cache:clear --no-warmup'" command:  
                                                                             
                                                                             
                                                                             
                                                                             
                                                                             
    [Doctrine\DBAL\Exception\DriverException]                                
    An exception occured in driver: could not find driver                    
                                                                             
                                                                             
                                                                             
    [Doctrine\DBAL\Driver\PDOException]                                      
    could not find driver                                                    
                                                                             
                                                                             
                                                                             
    [PDOException]                                                           
    could not find driver                                                    
                                                                             
                                                                             
  .                                                                          
                                                                             

install [--prefer-source] [--prefer-dist] [--dry-run] [--dev] [--no-dev] [--no-custom-installers] [--no-autoloader] [--no-scripts] [--no-progress] [--no-suggest] [-v|vv|vvv|--verbose] [-o|--optimize-autoloader] [-a|--classmap-authoritative] [--apcu-autoloader] [--ignore-platform-reqs] [--] [<packages>]...



[vagrant@kubedev-172-17-4-59 edusoho]$ php app/console server:run 0.0.0.0:8080

                                                         
  [Doctrine\DBAL\Exception\DriverException]              
  An exception occured in driver: could not find driver  
                                                         

                                       
  [Doctrine\DBAL\Driver\PDOException]  
  could not find driver                
                                       

                         
  [PDOException]         
  could not find driver  
                         

