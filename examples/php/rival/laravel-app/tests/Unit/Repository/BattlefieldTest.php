<?php
namespace Tests\Unit\Repository;

use Tests\TestCase;
// use Illuminate\Foundation\Testing\DatabaseMigrations;
// use Illuminate\Foundation\Testing\DatabaseTransactions;

// require the Faker autoloader
// require_once __DIR__.'/../../../vendor/fzaninotto/faker/src/autoload.php';
use Faker\Factory /* as FakerFactory*/ ;

use App\Battlefield;
use App\Repository\BattlefieldRepository;

class BattlefieldTest extends TestCase
{
    // use DatabaseMigrations;

    /**
     * A creation test.
     *
     * @return void
     */
    public function testBattlefieldModel()
    {
        $bf1 = factory(\App\Battlefield::class)->make();
        
        $this->assertNotNull($bf1->save());
        echo $bf1->id;
        
        $bf2 = factory(\App\Battlefield::class)->states('ready')->create();

    }
    
    /**
     * A creation test.
     *
     * @return void
     */
    public function testBattlefieldRepo()
    {
        $repo = new BattlefieldRepository;

        $faker = \Faker\Factory::create();
        // $faker = FakerFactory::create();
        
        $id = $repo->create($faker->name, $faker->realText(20), $faker->word);
        $this->assertNotNull($id);
        echo $id;
        
        $id = $repo->insert($faker->name, $faker->realText(20), $faker->word);
        $this->assertNotNull($id);

    }
    
}
