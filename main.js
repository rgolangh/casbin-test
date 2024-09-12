async function main() {
    const { newEnforcer } = require('casbin');
    // For browser:
    //import { newEnforcer } from 'casbin';

    const enforcer = await newEnforcer('model.conf', 'policy.csv');

    console.log(await enforcer.enforce("user:default/guest", "orchestrator.workflows.foo", "read"));
}


main();
