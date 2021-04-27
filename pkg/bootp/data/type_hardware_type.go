package data

const (
	HardwareType0_Unknown                       HardwareType = 0
	HardwareType1_Ethernet                      HardwareType = 1
	HardwareType2_ExperimentalEthernet          HardwareType = 2
	HardwareType3_AmateurRadioA_X25             HardwareType = 3
	HardwareType4_ProteonProNetTokenRing        HardwareType = 4
	HardwareType5_Chaos                         HardwareType = 5
	HardwareType6_Ieee802Networks               HardwareType = 6
	HardwareType7_ARCNET                        HardwareType = 7
	HardwareType8_HyperChannel                  HardwareType = 8
	HardwareType9_LanStar                       HardwareType = 9
	HardwareType10_AutonetShortAddress          HardwareType = 10
	HardwareType11_LocalTalk                    HardwareType = 11
	HardwareType12_LocalNet                     HardwareType = 12
	HardwareType13_UltraLink                    HardwareType = 13
	HardwareType14_SMDS                         HardwareType = 14
	HardwareType15_FrameRelay                   HardwareType = 15
	HardwareType16_AsynchronousTransmissionMode HardwareType = 16
)

const (
	descriptionHardwareType1_Unknown                       = "Unknown"
	descriptionHardwareType1_Ethernet                      = "Ethernet (10Mb)"
	descriptionHardwareType2_ExperimentalEthernet          = "Experimental Ethernet (3Mb)"
	descriptionHardwareType3_AmateurRadioA_X25             = "Amateur Radio AX.25"
	descriptionHardwareType4_ProteonProNetTokenRing        = "Proteon ProNET Token Ring"
	descriptionHardwareType5_Chaos                         = "Chaos"
	descriptionHardwareType6_Ieee802Networks               = "IEEE 802 Networks"
	descriptionHardwareType7_ARCNET                        = "ARCNET"
	descriptionHardwareType8_HyperChannel                  = "Hyperchannel"
	descriptionHardwareType9_LanStar                       = "Lanstar"
	descriptionHardwareType10_AutonetShortAddress          = "Autonet Short Address"
	descriptionHardwareType11_LocalTalk                    = "LocalTalk"
	descriptionHardwareType12_LocalNet                     = "LocalNet (IBM PCNet or SYTEK LocalNET)"
	descriptionHardwareType13_UltraLink                    = "Ultra link"
	descriptionHardwareType14_SMDS                         = "SMDS"
	descriptionHardwareType15_FrameRelay                   = "Frame Relay"
	descriptionHardwareType16_AsynchronousTransmissionMode = "Asynchronous Transmission Mode (ATM)"
)

var hardwareTypeDescription = map[HardwareType]string{
	HardwareType0_Unknown:                       descriptionHardwareType1_Unknown,
	HardwareType1_Ethernet:                      descriptionHardwareType1_Ethernet,
	HardwareType2_ExperimentalEthernet:          descriptionHardwareType2_ExperimentalEthernet,
	HardwareType3_AmateurRadioA_X25:             descriptionHardwareType3_AmateurRadioA_X25,
	HardwareType4_ProteonProNetTokenRing:        descriptionHardwareType4_ProteonProNetTokenRing,
	HardwareType5_Chaos:                         descriptionHardwareType5_Chaos,
	HardwareType6_Ieee802Networks:               descriptionHardwareType6_Ieee802Networks,
	HardwareType7_ARCNET:                        descriptionHardwareType7_ARCNET,
	HardwareType8_HyperChannel:                  descriptionHardwareType8_HyperChannel,
	HardwareType9_LanStar:                       descriptionHardwareType9_LanStar,
	HardwareType10_AutonetShortAddress:          descriptionHardwareType10_AutonetShortAddress,
	HardwareType11_LocalTalk:                    descriptionHardwareType11_LocalTalk,
	HardwareType12_LocalNet:                     descriptionHardwareType12_LocalNet,
	HardwareType13_UltraLink:                    descriptionHardwareType13_UltraLink,
	HardwareType14_SMDS:                         descriptionHardwareType14_SMDS,
	HardwareType15_FrameRelay:                   descriptionHardwareType15_FrameRelay,
	HardwareType16_AsynchronousTransmissionMode: descriptionHardwareType16_AsynchronousTransmissionMode,
}

// HardwareType this is type of hardware address
// See: RFC 951 section 3
type HardwareType int8

// Description return description of HardwareType
func (h HardwareType) Description() string {
	desc, exists := hardwareTypeDescription[h]
	if !exists {
		return hardwareTypeDescription[HardwareType0_Unknown]
	}

	return desc
}
