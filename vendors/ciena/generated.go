// Code generated by radius-dict-gen. DO NOT EDIT.

package ciena

import (
	"strconv"

	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_Ciena_VendorID = 1271
)

func _Ciena_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_Ciena_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _Ciena_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Ciena_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				values = append(values, vsa[2:int(vsaLen)])
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _Ciena_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _Ciena_VendorID {
			continue
		}
		for len(vsa) >= 3 {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				break
			}
			if vsaTyp == typ {
				return vsa[2:int(vsaLen)], true
			}
			vsa = vsa[int(vsaLen):]
		}
	}
	return
}

func _Ciena_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Ciena_VendorID {
			i++
			continue
		}
		for j := 0; len(vsa[j:]) >= 3; {
			vsaTyp, vsaLen := vsa[0], vsa[1]
			if int(vsaLen) > len(vsa[j:]) || vsaLen < 3 {
				i++
				break
			}
			if vsaTyp == typ {
				vsa = append(vsa[:j], vsa[j+int(vsaLen):]...)
			}
			j += int(vsaLen)
		}
		if len(vsa) > 0 {
			copy(avp.Attribute[4:], vsa)
			i++
		} else {
			p.Attributes = append(p.Attributes[:i], p.Attributes[i+i:]...)
		}
	}
	return _Ciena_AddVendor(p, typ, attr)
}

func _Ciena_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _Ciena_VendorID {
			i++
			continue
		}
		offset := 0
		for len(vsa[offset:]) >= 3 {
			vsaTyp, vsaLen := vsa[offset], vsa[offset+1]
			if int(vsaLen) > len(vsa) || vsaLen < 3 {
				continue vsaLoop
			}
			if vsaTyp == typ {
				copy(vsa[offset:], vsa[offset+int(vsaLen):])
				vsa = vsa[:len(vsa)-int(vsaLen)]
			} else {
				offset += int(vsaLen)
			}
		}
		if offset == 0 {
			p.Attributes = append(p.Attributes[:i], p.Attributes[i+1:]...)
		} else {
			i++
		}
	}
	return
}

func Ciena1CRole_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Ciena_AddVendor(p, 200, a)
}

func Ciena1CRole_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Ciena_AddVendor(p, 200, a)
}

func Ciena1CRole_Get(p *radius.Packet) (value []byte) {
	value, _ = Ciena1CRole_Lookup(p)
	return
}

func Ciena1CRole_GetString(p *radius.Packet) (value string) {
	value, _ = Ciena1CRole_LookupString(p)
	return
}

func Ciena1CRole_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Ciena_GetsVendor(p, 200) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func Ciena1CRole_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Ciena_GetsVendor(p, 200) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func Ciena1CRole_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Ciena_LookupVendor(p, 200)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func Ciena1CRole_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Ciena_LookupVendor(p, 200)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func Ciena1CRole_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Ciena_SetVendor(p, 200, a)
}

func Ciena1CRole_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Ciena_SetVendor(p, 200, a)
}

func Ciena1CRole_Del(p *radius.Packet) {
	_Ciena_DelVendor(p, 200)
}

func CienaBPMCPRole_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Ciena_AddVendor(p, 220, a)
}

func CienaBPMCPRole_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Ciena_AddVendor(p, 220, a)
}

func CienaBPMCPRole_Get(p *radius.Packet) (value []byte) {
	value, _ = CienaBPMCPRole_Lookup(p)
	return
}

func CienaBPMCPRole_GetString(p *radius.Packet) (value string) {
	value, _ = CienaBPMCPRole_LookupString(p)
	return
}

func CienaBPMCPRole_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _Ciena_GetsVendor(p, 220) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CienaBPMCPRole_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _Ciena_GetsVendor(p, 220) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CienaBPMCPRole_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _Ciena_LookupVendor(p, 220)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func CienaBPMCPRole_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _Ciena_LookupVendor(p, 220)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func CienaBPMCPRole_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _Ciena_SetVendor(p, 220, a)
}

func CienaBPMCPRole_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _Ciena_SetVendor(p, 220, a)
}

func CienaBPMCPRole_Del(p *radius.Packet) {
	_Ciena_DelVendor(p, 220)
}

type CienaWSUserPriv uint32

const (
	CienaWSUserPriv_Value_WaveserverPrivLimited CienaWSUserPriv = 1
	CienaWSUserPriv_Value_WaveserverPrivAdmin   CienaWSUserPriv = 2
	CienaWSUserPriv_Value_WaveserverPrivSuper   CienaWSUserPriv = 3
	CienaWSUserPriv_Value_WaveserverPrivDiag    CienaWSUserPriv = 4
)

var CienaWSUserPriv_Strings = map[CienaWSUserPriv]string{
	CienaWSUserPriv_Value_WaveserverPrivLimited: "Waveserver-Priv-Limited",
	CienaWSUserPriv_Value_WaveserverPrivAdmin:   "Waveserver-Priv-Admin",
	CienaWSUserPriv_Value_WaveserverPrivSuper:   "Waveserver-Priv-Super",
	CienaWSUserPriv_Value_WaveserverPrivDiag:    "Waveserver-Priv-Diag",
}

func (a CienaWSUserPriv) String() string {
	if str, ok := CienaWSUserPriv_Strings[a]; ok {
		return str
	}
	return "CienaWSUserPriv(" + strconv.FormatUint(uint64(a), 10) + ")"
}

func CienaWSUserPriv_Add(p *radius.Packet, value CienaWSUserPriv) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Ciena_AddVendor(p, 10, a)
}

func CienaWSUserPriv_Get(p *radius.Packet) (value CienaWSUserPriv) {
	value, _ = CienaWSUserPriv_Lookup(p)
	return
}

func CienaWSUserPriv_Gets(p *radius.Packet) (values []CienaWSUserPriv, err error) {
	var i uint32
	for _, attr := range _Ciena_GetsVendor(p, 10) {
		i, err = radius.Integer(attr)
		if err != nil {
			return
		}
		values = append(values, CienaWSUserPriv(i))
	}
	return
}

func CienaWSUserPriv_Lookup(p *radius.Packet) (value CienaWSUserPriv, err error) {
	a, ok := _Ciena_LookupVendor(p, 10)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	var i uint32
	i, err = radius.Integer(a)
	if err != nil {
		return
	}
	value = CienaWSUserPriv(i)
	return
}

func CienaWSUserPriv_Set(p *radius.Packet, value CienaWSUserPriv) (err error) {
	a := radius.NewInteger(uint32(value))
	return _Ciena_SetVendor(p, 10, a)
}

func CienaWSUserPriv_Del(p *radius.Packet) {
	_Ciena_DelVendor(p, 10)
}
