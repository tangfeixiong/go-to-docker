// [vagrant@bogon laravel-app]$ php artisan make:controller BattlefiledRepoController
// Controller created successfully.

<?php

namespace App\Http\Controllers\ReSTful;

use Illuminate\Http\Request;

use App\Repository\BattlefieldRepository;

class BattlefiledRepoController extends Controller
{

    /**
     * The user repository instance.
     */
    protected $bfrepo;

    /**
     * Create a new controller instance.
     *
     * @param  UserRepository  $users
     * @return void
     */
    public function __construct(BattlefieldRepository $repo)
    {
        $this->bfrepo = $repo;
    }
    
    /** POST /api/v1/battlefields
     *
     */
    public function store(Request $request)
    {
        return Article::create($request->all());
    }


}
