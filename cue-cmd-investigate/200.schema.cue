// https://cuelang.org/docs/usecases/datadef/
// Release notes:
// - You can now specify your age and your hobby!
#V1: {
    age:   >=0 & <=100
    hobby: string
}
// Release notes:
// - People get to be older than 100, so we relaxed it.
// - It seems not many people have a hobby, so we made it optional.
#V2: {
    age:    >=0 & <=150 // people get older now
    hobby?: string      // some people don't have a hobby
}
// Release notes:
// - Actually no one seems to have a hobby nowadays anymore, so we dropped the field.
#V3: {
    age: >=0 & <=150
}
