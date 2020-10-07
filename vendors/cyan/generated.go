// Code generated by radius-dict-gen. DO NOT EDIT.

package cyan

import (
	"layeh.com/radius"
	"layeh.com/radius/rfc2865"
)

const (
	_CyanInc_VendorID = 28533
)

func _CyanInc_AddVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	var vsa radius.Attribute
	vendor := make(radius.Attribute, 2+len(attr))
	vendor[0] = typ
	vendor[1] = byte(len(vendor))
	copy(vendor[2:], attr)
	vsa, err = radius.NewVendorSpecific(_CyanInc_VendorID, vendor)
	if err != nil {
		return
	}
	p.Add(rfc2865.VendorSpecific_Type, vsa)
	return
}

func _CyanInc_GetsVendor(p *radius.Packet, typ byte) (values []radius.Attribute) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _CyanInc_VendorID {
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

func _CyanInc_LookupVendor(p *radius.Packet, typ byte) (attr radius.Attribute, ok bool) {
	for _, avp := range p.Attributes {
		if avp.Type != rfc2865.VendorSpecific_Type {
			continue
		}
		attr := avp.Attribute
		vendorID, vsa, err := radius.VendorSpecific(attr)
		if err != nil || vendorID != _CyanInc_VendorID {
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

func _CyanInc_SetVendor(p *radius.Packet, typ byte, attr radius.Attribute) (err error) {
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _CyanInc_VendorID {
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
	return _CyanInc_AddVendor(p, typ, attr)
}

func _CyanInc_DelVendor(p *radius.Packet, typ byte) {
vsaLoop:
	for i := 0; i < len(p.Attributes); {
		avp := p.Attributes[i]
		if avp.Type != rfc2865.VendorSpecific_Type {
			i++
			continue
		}
		vendorID, vsa, err := radius.VendorSpecific(avp.Attribute)
		if err != nil || vendorID != _CyanInc_VendorID {
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

func CyanIncUserRoles_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _CyanInc_AddVendor(p, 1, a)
}

func CyanIncUserRoles_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _CyanInc_AddVendor(p, 1, a)
}

func CyanIncUserRoles_Get(p *radius.Packet) (value []byte) {
	value, _ = CyanIncUserRoles_Lookup(p)
	return
}

func CyanIncUserRoles_GetString(p *radius.Packet) (value string) {
	value, _ = CyanIncUserRoles_LookupString(p)
	return
}

func CyanIncUserRoles_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _CyanInc_GetsVendor(p, 1) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CyanIncUserRoles_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _CyanInc_GetsVendor(p, 1) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CyanIncUserRoles_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _CyanInc_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func CyanIncUserRoles_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _CyanInc_LookupVendor(p, 1)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func CyanIncUserRoles_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _CyanInc_SetVendor(p, 1, a)
}

func CyanIncUserRoles_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _CyanInc_SetVendor(p, 1, a)
}

func CyanIncUserRoles_Del(p *radius.Packet) {
	_CyanInc_DelVendor(p, 1)
}

func CyanIncAcctEventText_Add(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _CyanInc_AddVendor(p, 100, a)
}

func CyanIncAcctEventText_AddString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _CyanInc_AddVendor(p, 100, a)
}

func CyanIncAcctEventText_Get(p *radius.Packet) (value []byte) {
	value, _ = CyanIncAcctEventText_Lookup(p)
	return
}

func CyanIncAcctEventText_GetString(p *radius.Packet) (value string) {
	value, _ = CyanIncAcctEventText_LookupString(p)
	return
}

func CyanIncAcctEventText_Gets(p *radius.Packet) (values [][]byte, err error) {
	var i []byte
	for _, attr := range _CyanInc_GetsVendor(p, 100) {
		i = radius.Bytes(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CyanIncAcctEventText_GetStrings(p *radius.Packet) (values []string, err error) {
	var i string
	for _, attr := range _CyanInc_GetsVendor(p, 100) {
		i = radius.String(attr)
		if err != nil {
			return
		}
		values = append(values, i)
	}
	return
}

func CyanIncAcctEventText_Lookup(p *radius.Packet) (value []byte, err error) {
	a, ok := _CyanInc_LookupVendor(p, 100)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.Bytes(a)
	return
}

func CyanIncAcctEventText_LookupString(p *radius.Packet) (value string, err error) {
	a, ok := _CyanInc_LookupVendor(p, 100)
	if !ok {
		err = radius.ErrNoAttribute
		return
	}
	value = radius.String(a)
	return
}

func CyanIncAcctEventText_Set(p *radius.Packet, value []byte) (err error) {
	var a radius.Attribute
	a, err = radius.NewBytes(value)
	if err != nil {
		return
	}
	return _CyanInc_SetVendor(p, 100, a)
}

func CyanIncAcctEventText_SetString(p *radius.Packet, value string) (err error) {
	var a radius.Attribute
	a, err = radius.NewString(value)
	if err != nil {
		return
	}
	return _CyanInc_SetVendor(p, 100, a)
}

func CyanIncAcctEventText_Del(p *radius.Packet) {
	_CyanInc_DelVendor(p, 100)
}