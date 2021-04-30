use anchor_lang::prelude::*;

// Enforces a single bumping index number.
#[derive(AnchorSerialize, AnchorDeserialize, Clone, Copy, Debug)]
pub struct Index(pub u32);

#[repr(C)]
#[derive(AnchorSerialize, AnchorDeserialize, Clone, Copy, Debug)]
pub struct BridgeConfig {
    /// Period for how long a guardian set is valid after it has been replaced by a new one.  This
    /// guarantees that VAAs issued by that set can still be submitted for a certain period.  In
    /// this period we still trust the old guardian set.
    pub guardian_set_expiration_time: u32,
}
