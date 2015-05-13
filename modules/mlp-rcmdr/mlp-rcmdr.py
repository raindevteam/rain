import string
import math
import random
 
    ###########################################################
    ########## Which episode of MLP should I watch? ###########
    ###########################################################
    """season1"""
    s101 = "s1e01 - Friendship is Magic Part 1"
    s102 = "s1e02 - Friendship is Magic Part 2"
    s103 = "s1e03 - The Ticket Master"
    s104 = "s1e04 - Applebuck Season"
    s105 = "s1e05 - Griffon the Brush-off"
    s106 = "s1e06 - Boast Busters"
    s107 = "s1e07 - Dragonshy"
    s108 = "s1e08 - Look Before you Sleep"
    s109 = "s1e09 - Bridle Gossip"
    s110 = "s1e10 - Swarm of the Century"
    s111 = "s1e11 - Winter Wrap Up"
    s112 = "s1e12 - Call of the Cutie"
    s113 = "s1e13 - Fall Weather Friends"
    s114 = "s1e14 - Suited for Success"
    s115 = "s1e15 - Feeling Pinkie Keen"
    s116 = "s1e16 - Sonic Rainboom"
    s117 = "s1e17 - Stare Master"
    s118 = "s1e18 - The Show Stoppers"
    s119 = "s1e19 - A Dog and Pony Show"
    s120 = "s1e20 - Green Isn't Your Color"
    s121 = "s1e21 - Over a Barrel"
    s122 = "s1e22 - A Bird in the Hoof"
    s123 = "s1e23 - The Cutie Mark Chronicles"
    s124 = "s1e24 - Owls Well That Ends Well"
    s125 = "s1e25 - Party of One"
    s126 = "s1e26 - The Best Night Ever"
    """season2"""
    s201 = "s2e01 - Return of Harmony Part 1"
    s202 = "s2e02 - Return of Harmony Part 2"
    s203 = "s2e03 - Lesson Zero"
    s204 = "s2e04 - Luna Eclipsed"
    s205 = "s2e05 - Sisterhooves Social"
    s206 = "s2e06 - The Cutie Pox"
    s207 = "s2e07 - May the Best Pet Win"
    s208 = "s2e08 - The Mysterious Mare Do Well"
    s209 = "s2e09 - Sweet and Elite"
    s210 = "s2e10 - Secret of my Excess"
    s211 = "s2e11 - Hearths Warming Eve"
    s212 = "s2e12 - Family Appreciation Day"
    s213 = "s2e13 - Baby Cakes"
    s214 = "s2e14 - The Last Roundup"
    s215 = "s2e15 - The Super Speedy Cider Squeezy 6000"
    s216 = "s2e16 - Read it and Weep"
    s217 = "s2e17 - Hearts and Hooves Day"
    s218 = "s2e18 - A Friend in Deed"
    s219 = "s2e19 - Putting Your Hoof Down"
    s220 = "s2e20 - It's About Time"
    s221 = "s2e21 - Dragon Quest"
    s222 = "s2e22 - Hurricane Fluttershy"
    s223 = "s2e23 - Ponyville Confidential"
    s224 = "s2e24 - MMMystery on the Friendship Express"
    s225 = "s2e25 - A canterlot Wedding Part 1"
    s226 = "s2e26 - A Canterlot Wedding Part 2"
    """season3"""
    s301 = "s3e01 - The Crystal Empire Part 1"
    s302 = "s3e02 - The Crystal Empire Part 2"
    s303 = "s3e03 - Too Many Pinkie Pies"
    s304 = "s3e04 - One Bad Apple"
    s305 = "s3e05 - Magic Duel"
    s306 = "s3e06 - Sleepless in Ponyville"
    s307 = "s3e07 - Wonderbolt Academy"
    s308 = "s3e08 - Apple Family Reunion"
    s309 = "s3e09 - Spike at Your Service"
    s310 = "s3e10 - Keep Calm and Flutter On"
    s311 = "s3e11 - Just For Sidekicks"
    s312 = "s3e12 - Games Ponies Play"
    s313 = "s3e13 - Magical Mystery Cure"
    """season4"""
    s401 = "s4e01 - Princess Twilight Part 1"
    s402 = "s4e02 - Princess Twilight Part 2"
    s403 = "s4e03 - Castle-Mania"
    s404 = "s4e04 - Daring Don't"
    s405 = "s4e05 - Flight to the Finish"
    s406 = "s4e06 - Power Ponies"
    s407 = "s4e07 - Bats!"
    s408 = "s4e08 - Rarity Takes Manehattan"
    s409 = "s4e09 - Pinkie Apple Pie"
    s410 = "s4e10 - Rainbow Falls"
    s411 = "s4e11 - Three's a Crowd"
    s412 = "s4e12 - Pinkie Pride"
    s413 = "s4e13 - Simple Ways"
    s414 = "s4e14 - Filli Vanilli"
    s415 = "s4e15 - Twilight Time"
    s416 = "s4e16 - It Ain't Easy Being Breezies"
    s417 = "s4e17 - Somepony to Watch Over Me"
    s418 = "s4e18 - Maud Pie"
    s419 = "s4e19 - For Whom the Sweetie Belle Toils"
    s420 = "s4e20 - Leap of Faith"
    s421 = "s4e21 - Testing, Testing,1,2,3"
    s422 = "s4e22 - "
    s423 = "s4e23 - "
    s424 = "s4e24 - "
    s425 = "s4e25 - Twilight's Kingdom Part 1"
    s426 = "s4e26 - Twilight's Kingdom Part 2"
    ###########################################################
    season1List = [s101, s102, s103, s104, s105, s106, s107, s108, s109, s110, s111, s112, s113, s114, s115, s116, s117, s118, s119, s120, s121, s122, s123, s124, s125, s126]
    season2List = [s201, s202, s203, s204, s205, s206, s207, s208, s209, s210, s211, s212, s213, s214, s215, s216, s217, s218, s219, s220, s221, s222, s223, s224, s225, s226]
    season3List = [s301, s302, s303, s304, s305, s306, s307, s308, s309, s310, s311, s312, s313]
    season4List = [s401, s402, s403, s404, s405, s406, s407, s408, s409, s410, s411, s412, s413, s414, s415, s416, s417, s418, s419, s420, s421, s422, s423, s424, s425, s426]
    fullList = [s101, s102, s103, s104, s105, s106, s107, s108, s109, s110, s111, s112, s113, s114, s115, s116, s117, s118, s119, s120, s121, s122, s123, s124, s125, s126, s201, s202, s203, s204, s205, s206, s207, s208, s209, s210, s211, s212, s213, s214, s215, s216, s217, s218, s219, s220, s221, s222, s223, s224, s225, s226, s301, s302, s303, s304, s305, s306, s307, s308, s309, s310, s311, s312, s313, s401, s402, s403, s404, s405, s406, s407, s408, s409, s410, s411, s412, s413, s414, s415, s416, s417, s418, s419, s420, s421, s422, s423, s424, s425, s426]
 
    response = " "
    while response != "n":
        season = " "
        response = " "
        while season != "0" and season != "1" and season != "2" and season != "3" and season != "4":
            season = seasonP
            if season == "1":
                episode = random.choice(season1List)
            elif season == "2":
                episode = random.choice(seasonList)
            elif season == "3":
                episode = random.choice(season3List)
            elif season == "4":
                episode = random.choice(season4List)
            elif season == "0":
                episode = random.choice(fullList)
            else:
                print "Invalid input, try again."
 
        print "You should watch " + episode
        print "Choose again? y/n"
       
        while response != "n" and response != "y":
            response = raw_input()
            if response == "n":
                print "Have fun!"
            elif response == "y":
                print "Choosing again..."
            else:
                print "Invalid input, try again"
                response = " "