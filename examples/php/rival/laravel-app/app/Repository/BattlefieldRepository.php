<?php

namespace App\Repository;

use Illuminate\Support\Facades\DB;
use App\Battlefield;

class BattlefieldRepository
{
    
    /**
     * Create a new 'Battlefield' instance.
     *
     * @param  Request  $request
     * @return Response
     */
    public function create($name, $desc, $matchItem, $status='not-ready', $payment='beta')
    {
        // Validate the ...

        $bf = new Battlefield;

        $bf->fill([
            'name' => $name,
            'description' => $desc,
            'match_item' => $matchItem,
            'status' => $status,
            'payment' => $payment
        ]);
        
        if ($bf->save())
            return $bf->id;
    }
    
    public function insert($name, $desc, $matchItem, $status='not-ready', $payment='beta')
    {
        $id = DB::table('battlefield')->insert([
            [
                'name' => $name,
                'description' => $desc,
                'match_item' => $matchItem,
                'status' => $status,
                'payment' => $payment
            ],
        ]);        
        return $id;
    }
    
    /**
     * Show a list of all of 'Battlefield'.
     *
     * @return Response
     */
    public function index()
    {
        $bf = DB::table('battlefield')->get();

        // return view('battlefield.index', ['battlefield' => $bf]);
        return $bf;
    }
}