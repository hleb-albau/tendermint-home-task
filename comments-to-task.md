# Comments

Honestly, given task paralyzed me at first. It's not oblivious, what functionality expected from the module. There are
two confusing things from my point of view: namings, and task idea in whole.

1. There is line in the **Task** section - is to be a simple registry for **chain names**. You expect some sort of
   **name** auction here, similar to DNS one.
2. Next, in the **Code** section we create a **Chain** entity. And in **Further thoughts on the implementation** section
   there is a link to some old **spn** proposal, where we expand that **Chain** objects with other fields. At this point
   **Chain** becomes your module main entity. So you have a **Chain** with a properties, that can be used by other tools
   to implement their functionality (such as **spn**). Chain name itself just an identifier. No more. So is it registry
   for chain names or registry for chains?
3. I
   found [spn module draft](https://github.com/tendermint/spn/blob/c0c163e77c96b137f1656f9c4f51f96569614264/docs/cns.md)
   confusing itself. It uses `ChainName` instead of simple `Chain` through all spec, starting from the title. As for
   current [main branch code](https://github.com/tendermint/spn/blob/develop/x/launch/types/chain.go) it is already
   renamed to `Chain`. Btw, a like an `Offer` mechanics from proposal.
4. In general, tasks seems like create an entity with CRUD+(Transfer)(Fund) functionality, without context, how such
   module will be used. Then, when you apply in to `Chain` entity, there a lot of questions arise:
    - how we can prove ownership of chain? every one can registry chain with any name? What if someone grab your name
      for an existing chain? Who should be and owner of highly decentralize-developed chains, such as bitcoin? Such
      module make sense in context of **spn**, where launching **new** chains occurs, but there is no word about it, and
      when you try to think about this task generally, it blows your mind.

So, I decided to keep it simple :) skipping auctions, offers and other mechanisms.